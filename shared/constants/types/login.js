// @flow
import HiddenString from '../../util/hidden-string'
export type Mode =
  | 'codePageModeScanCode'
  | 'codePageModeShowCode'
  | 'codePageModeEnterText'
  | 'codePageModeShowText'

export type DeviceRole =
  | 'codePageDeviceRoleExistingPhone'
  | 'codePageDeviceRoleNewPhone'
  | 'codePageDeviceRoleExistingComputer'
  | 'codePageDeviceRoleNewComputer'

// It's the b64 encoded value used to render the image
export type QRCode = HiddenString
export type State = {
  codePageCameraBrokenMode: boolean,
  codePageCodeCountDown: number,
  codePageEnterCodeErrorText: string,
  codePageMode: ?Mode,
  codePageMyDeviceRole: ?DeviceRole,
  codePageOtherDeviceRole: ?DeviceRole,
  codePageQrCode: ?QRCode,
  codePageQrCodeScanned: boolean,
  codePageQrScanned: ?QRCode,
  codePageTextCode: ?HiddenString,
  configuredAccounts: ?Array<{|hasStoredSecret: boolean, username: string|}>,
  forgotPasswordError: ?Error,
  forgotPasswordSubmitting: boolean,
  forgotPasswordSuccess: boolean,
  justDeletedSelf: ?string,
  justRevokedSelf: ?string,
  loginError: ?string,
  registerUserPassError: ?string,
  registerUserPassLoading: boolean,
  waitingForResponse: boolean,
}
