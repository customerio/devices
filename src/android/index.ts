import devices from '../data/android.json';
import { AndroidDevice as Device } from '../types';

function deviceList(): Device[] {
  return devices;
}

/**
 * Get the first matching device's marketing name from the list of devices.
 * If no device is found, returns the given model name.
 *
 * @param model
 */
export const getAndroidMarketingName = (model: string): string => {
  const device = deviceList().find((device) => device.model === model);

  let marketingName = device?.marketingName;

  if (!marketingName) {
    marketingName = model;
  }

  return marketingName;
};
