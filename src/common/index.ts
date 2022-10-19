import { getIosDevices } from '../ios';
import { getAndroidDevices } from '../android';

/**
 * Get marketing name for a given device model or identifier.
 * If no device is found, returns null
 *
 * @param deviceModel - case insensitive
 */
export const getMarketingName = (deviceModel: string): string | null => {
  const iosDevices = getIosDevices(deviceModel);

  if (iosDevices.length > 0) {
    return iosDevices[0].generation;
  }

  const androidDevices = getAndroidDevices(deviceModel);
  if (androidDevices.length > 0) {
    return androidDevices[0].marketingName;
  }

  return null;
};
