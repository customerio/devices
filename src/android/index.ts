import devicesArray from '../data/android.json';
import { AndroidDevice as Device } from '../types';

function deviceList(): Device[] {
  const devices = [];
  devicesArray.forEach((deviceItem) => {
    devices.push({
      retailBranding: deviceItem[0],
      marketingName: deviceItem[1],
      device: deviceItem[2],
      model: deviceItem[3],
    });
  });
  return devices;
}

/**
 * Get all possible android devices for a given device model.
 */
export const getAndroidDevices = (model: string): Device[] => {
  return deviceList().filter(
    (device) => device.model.toLowerCase() === model.toLowerCase(),
  );
};

export const getAndroidDevicesFromMarketingName = (
  marketingName: string,
): Device[] => {
  return deviceList().filter(
    (device) =>
      device.marketingName.toLowerCase() === marketingName.toLowerCase(),
  );
};
