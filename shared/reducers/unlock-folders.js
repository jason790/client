// @flow
import * as I from 'immutable'
import * as UnlockFoldersGen from '../actions/unlock-folders-gen'
import * as Constants from '../constants/unlock-folders'
import * as Types from '../constants/types/unlock-folders'
import {toDeviceType} from '../constants/devices'

const initialState = Constants.makeState()

export default function(state: Types.State = initialState, action: UnlockFoldersGen.Actions): Types.State {
  switch (action.type) {
    case UnlockFoldersGen.resetStore:
      return initialState
    case UnlockFoldersGen.closeDone:
      return state.set('popupOpen', false)
    case UnlockFoldersGen.waiting: {
      const {waiting} = action.payload
      return state.set('waiting', waiting)
    }
    case UnlockFoldersGen.onBackFromPaperKey:
      return state.set('paperkeyError', '').set('phase', 'promptOtherDevice')
    case UnlockFoldersGen.toPaperKeyInput:
      return state.set('phase', 'paperKeyInput')
    case UnlockFoldersGen.checkPaperKeyDone:
      if (action.error) {
        return state.set('paperkeyError', action.payload.error)
      }
      return state.set('phase', 'success')
    case UnlockFoldersGen.finish:
      return state.set('popupOpen', false).set('phase', 'dead')
    case UnlockFoldersGen.newRekeyPopup: {
      const devices = I.List(
        action.payload.devices.map(({name, type, deviceID}) =>
          Constants.makeDevice({
            deviceID,
            name,
            type: toDeviceType(type),
          })
        )
      )
      return state
        .set('popupOpen', !!devices.count())
        .set('devices', devices)
        .set('sessionID', action.payload.sessionID)
    }
    // Saga only actions
    case UnlockFoldersGen.checkPaperKey:
    case UnlockFoldersGen.closePopup:
    case UnlockFoldersGen.openPopup:
    case UnlockFoldersGen.registerRekeyListener:
      return state
    default:
      // eslint-disable-next-line no-unused-expressions
      (action: empty) // if you get a flow error here it means there's an action you claim to handle but didn't
      return state
  }
}
