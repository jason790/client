// @flow
import * as Types from '../constants/types/route-tree'
import * as Constants from '../constants/route-tree'
import * as I from 'immutable'
import * as Saga from '../util/saga'
import {getPath} from '../route-tree'

import type {RouteDefParams, Path, PropsPath} from '../route-tree'
import type {TypedAction} from '../constants/types/flux'
import type {TypedState} from '../constants/reducer'

const pathActionTransformer = (action, oldState) => {
  const prevPath = oldState.routeTree ? getPath(oldState.routeTree.routeState) : I.List()
  const path = Array.from(action.payload.path.map(p => (typeof p === 'string' ? p : p.selected)))
  const parentPath = action.payload.parentPath && Array.from(action.payload.parentPath)
  return {
    payload: {
      prevPath,
      path,
      parentPath,
    },
    type: action.type,
  }
}

export function pathSelector(state: TypedState, parentPath?: Path): I.List<string> {
  return getPath(state.routeTree.routeState, parentPath)
}

// Set (or update) the tree of route definitions. Dispatched at initialization
// time and when route definitions update through HMR.
export function setRouteDef(routeDef: RouteDefParams): Types.SetRouteDef {
  return {
    type: Constants.setRouteDef,
    payload: {routeDef},
  }
}

// Switch to a new path, restoring the subpath it was previously on. E.g.:
//
//   (starting on /settings/invites)
//   navigate({selected: 'folders'})
//   => /folders
//   switchTo('settings')
//   => /settings/invites
//
// If parentPath is provided, the path will be switched to relative to
// parentPath without navigating to it.
export function switchTo(path: Path, parentPath?: Path): Types.SwitchTo {
  return {
    type: Constants.switchTo,
    payload: {path, parentPath},
    logTransformer: pathActionTransformer,
  }
}

// Navigate to a new absolute path. E.g.:
//
//   navigateTo({selected: 'foo', props: {prop1: 'hello'}}, {selected: 'bar', props: {prop2: 'world'}})
//   => /foo?prop1=hello/bar?prop2=world
//
// You can also specify path names as strings. This will select the name
// without changing props:
//
//    (starting on /foo?prop1=hello/bar?prop2=world)
//    navigateTo('foo', {selected: 'bar'})
//    => /foo?prop1=hello/bar
//
// If parentPath is provided, the path will be navigated to relative to
// parentPath without navigating to it.
export function navigateTo(
  path: PropsPath<*>,
  parentPath?: ?Path,
  navigationSource: Types.NavigationSource = 'user'
): Types.NavigateTo {
  return {
    type: Constants.navigateTo,
    payload: {path, parentPath, navigationSource},
    logTransformer: pathActionTransformer,
  }
}

// Navigate to a path relative to the current path.
// If parentPath is provided, the path will be appended relative to parentPath
// without navigating to it.
export function navigateAppend(path: PropsPath<*>, parentPath?: Path): Types.NavigateAppend {
  return {
    type: Constants.navigateAppend,
    payload: {path, parentPath},
    logTransformer: pathActionTransformer,
  }
}

// Navigate one step up from the current path.
export function navigateUp(): Types.NavigateUp {
  return {
    type: Constants.navigateUp,
    payload: null,
  }
}

// Do a navigate action if the path is still what is expected
export function putActionIfOnPath<T: TypedAction<*, *, *>>(
  expectedPath: Path,
  otherAction: T,
  parentPath?: Path
): Types.PutActionIfOnPath<T> {
  return {
    type: Constants.putActionIfOnPath,
    payload: {expectedPath, otherAction, parentPath},
  }
}

// Update the state object of a route at a specified path.
export function setRouteState(
  path: Path,
  partialState: {} | ((oldState: I.Map<string, any>) => I.Map<string, any>)
): Types.SetRouteState {
  return {
    type: Constants.setRouteState,
    payload: {path, partialState},
    logTransformer: pathActionTransformer,
  }
}

// Reset the props and state for a subtree.
export function resetRoute(path: Path): Types.ResetRoute {
  return {
    type: Constants.resetRoute,
    payload: {path},
    logTransformer: pathActionTransformer,
  }
}

function* _putActionIfOnPath({payload: {otherAction, expectedPath, parentPath}}: Types.PutActionIfOnPath<*>) {
  const state: TypedState = yield Saga.select()
  const currentPath = pathSelector(state, parentPath)
  if (I.is(I.List(expectedPath), currentPath)) {
    yield Saga.put(otherAction)
  }
}

function* routeSaga(): any {
  yield Saga.safeTakeEvery('routeTree:putActionIfOnPath', _putActionIfOnPath)
}

export default routeSaga
