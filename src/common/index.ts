import { getIosDevices } from '../ios';
import { getAndroidDevices } from '../android';
import { OsType } from '../types';

/**
 * Get marketing name for a given device model or identifier.
 * If no device is found, returns the model or identifier.
 *
 * @param deviceModel
 * @param os
 */
export const getMarketingName = (deviceModel: string, os: OsType): string => {
  if (os === OsType.IOS) {
    const devices = getIosDevices(deviceModel);
    if (devices.length > 0) {
      return devices[0].generation;
    }
  } else {
    const devices = getAndroidDevices(deviceModel);
    if (devices.length > 0) {
      return devices[0].marketingName;
    }
  }
  return deviceModel;
};
