// @flow
import * as SettingsGen from '../../actions/settings-gen'
import Bootstrapable from '../../util/bootstrapable'
import Landing from '.'
import {connect, type TypedState, compose} from '../../util/container'
import {navigateAppend} from '../../actions/route-tree'

const mapStateToProps = (state: TypedState, ownProps: {}) => {
  const {emails} = state.settings.email
  let accountProps
  if (emails.length > 0) {
    accountProps = {
      email: emails[0].email,
      isVerified: emails[0].isVerified,
      onChangeEmail: () => console.log('todo'),
      onChangePassphrase: () => console.log('todo'),
    }
  }

  const {planBilling: {availablePlans, usage, plan, paymentInfo}} = state
  let planProps
  if (plan && usage) {
    const freeSpaceGB = plan.gigabytes - usage.gigabytes
    const freeSpacePercentage = freeSpaceGB / plan.gigabytes
    planProps = {
      onUpgrade: () => console.log('todo'),
      onDowngrade: () => console.log('todo'),
      onInfo: () => console.log('todo'),
      selectedLevel: plan.planLevel,
      freeSpace: freeSpaceGB + 'GB',
      freeSpacePercentage,
      lowSpaceWarning: false,
      paymentInfo,
      onChangePaymentInfo: () => console.log('todo'),
    }
  }

  // When enabling planProps, we should check both for bootstrapDone:
  // let bootstrapDone = accountProps && planProps
  let bootstrapDone = !!accountProps

  return {
    bootstrapDone: bootstrapDone,
    originalProps: {
      account: accountProps,
      plan: planProps,
      plans: availablePlans,
    },
  }
}

const mapDispatchToProps = (dispatch: (a: any) => void, ownProps: {}) => ({
  onBootstrap: () => {
    dispatch(SettingsGen.createLoadSettings())
  },
  onChangePassphrase: () => dispatch(navigateAppend(['changePassphrase'])),
  onChangeEmail: () => dispatch(navigateAppend(['changeEmail'])),
  onInfo: selectedLevel => dispatch(navigateAppend([{selected: 'changePlan', props: {selectedLevel}}])),
})

const mergeProps = (stateProps, dispatchProps, ownProps: {}) => {
  if (!stateProps.bootstrapDone) {
    return {
      ...stateProps,
      onBootstrap: dispatchProps.onBootstrap,
    }
  }

  return {
    ...stateProps,
    originalProps: {
      ...stateProps.originalProps,
      account: {
        ...stateProps.originalProps.account,
        onChangeEmail: dispatchProps.onChangeEmail,
        onChangePassphrase: dispatchProps.onChangePassphrase,
      },
      plan: {
        ...stateProps.originalProps.plan,
        onInfo: selectedLevel => {
          dispatchProps.onInfo(selectedLevel)
        },
      },
    },
  }
}
export default compose(connect(mapStateToProps, mapDispatchToProps, mergeProps), Bootstrapable)(Landing)
