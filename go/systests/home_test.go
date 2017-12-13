package systests

import (
	"github.com/keybase/client/go/client"
	keybase1 "github.com/keybase/client/go/protocol/keybase1"
	"github.com/stretchr/testify/require"
	context "golang.org/x/net/context"
	"testing"
	"time"
)

func getHome(t *testing.T, u *userPlusDevice, markViewed bool) keybase1.HomeScreen {
	g := u.tc.G
	cli, err := client.GetHomeClient(g)
	require.NoError(t, err)
	home, err := cli.HomeGetScreen(context.TODO(), keybase1.HomeGetScreenArg{MarkViewed: markViewed, NumFollowSuggestionsWanted: 10})
	require.NoError(t, err)
	return home
}

func getBadgeState(t *testing.T, u *userPlusDevice) keybase1.BadgeState {
	g := u.tc.G
	cli, err := client.GetBadgerClient(g)
	require.NoError(t, err)
	ret, err := cli.GetBadgeState(context.TODO())
	require.NoError(t, err)
	return ret
}

func assertTodoPresent(t *testing.T, home keybase1.HomeScreen, wanted keybase1.HomeScreenTodoType, isBadged bool) {
	for _, item := range home.Items {
		typ, err := item.Data.T()
		if err != nil {
			t.Fatal(err)
		}
		if typ == keybase1.HomeScreenItemType_TODO {
			todo := item.Data.Todo()
			typ, err := todo.T()
			if err != nil {
				t.Fatal(err)
			}
			if typ == wanted {
				require.Equal(t, item.Badged, isBadged)
				return
			}
		}
	}
	t.Fatalf("Failed to find type %s in %+v", wanted, home)
}

func assertTodoNotPresent(t *testing.T, home keybase1.HomeScreen, wanted keybase1.HomeScreenTodoType) {
	for _, item := range home.Items {
		typ, err := item.Data.T()
		if err != nil {
			t.Fatal(err)
		}
		if typ == keybase1.HomeScreenItemType_TODO {
			todo := item.Data.Todo()
			typ, err := todo.T()
			if err != nil {
				t.Fatal(err)
			}
			if typ == wanted {
				t.Fatalf("Found type %s in %+v, but didn't want to ", wanted, home)
			}
		}
	}
}

func postBio(t *testing.T, u *userPlusDevice) {
	g := u.tc.G
	cli, err := client.GetUserClient(g)
	require.NoError(t, err)
	arg := keybase1.ProfileEditArg{
		FullName: "Boaty McBoatface",
		Location: "The Sea, The Sea",
		Bio:      "Just your average stupidly named vessel",
	}
	err = cli.ProfileEdit(context.TODO(), arg)
	require.NoError(t, err)
}

func TestHome(t *testing.T) {
	tt := newTeamTester(t)
	defer tt.cleanup()

	// It's important that we don't add with a paper key, because if we did, we'd
	// burn through the first two todo items (1. Device; 2. Paper Key), and then
	// we wouldn't get a badge for another day or so. So this way, we're only going
	// to have one item done (device), and two items will be in the next todo
	// list, with one badged. This is a bit brittle since if the server-side logic
	// changes, this test will break. But let's leave it for now.
	tt.addUserNoPaper("alice")
	alice := tt.users[0]
	g := alice.tc.G

	home := getHome(t, alice, true)
	initialVersion := home.Version

	require.True(t, (initialVersion > 0), "initial version should be > 0")
	assertTodoPresent(t, home, keybase1.HomeScreenTodoType_BIO, true)

	// Wait for a gregor message to fill in the badge state, for at most ~10s.
	// Hopefully this is enough for slow CI but you never know.
	pollForTrue := func(poller func(i int) bool) {
		// Hopefully this is enough for slow CI but you never know.
		wait := 10 * time.Millisecond
		found := false
		for i := 0; i < 10; i++ {
			if poller(i) {
				found = true
				break
			}
			g.Log.Debug("Didn't get an update; waiting %s more", wait)
			time.Sleep(wait)
			wait = wait * 2
		}
		require.True(t, found, "found result after poll")
	}

	var countPre int
	pollForTrue(func(i int) bool {
		badges := getBadgeState(t, alice)
		g.Log.Debug("Iter loop %d badge state: %+v", i, badges)
		countPre = badges.HomeTodoItems
		return (countPre == 1)
	})

	postBio(t, alice)

	pollForTrue(func(i int) bool {
		home = getHome(t, alice, true)
		badges := getBadgeState(t, alice)
		g.Log.Debug("Iter %d of check loop: Home is: %+v; BadgeState is: %+v", i, home, badges)
		return (home.Version > initialVersion && badges.HomeTodoItems < countPre)
	})

	assertTodoNotPresent(t, home, keybase1.HomeScreenTodoType_BIO)
}
