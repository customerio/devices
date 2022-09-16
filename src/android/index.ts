import devices from '../data/android.json';
import { AndroidDevice as Device } from '../types';

function deviceList(): Device[] {
  return devices;
}

/**
 * Get all possible android devices for a given device model.
 */
export const getAndroidDevices = (model: string): Device[] => {
  return deviceList().filter((device) => device.model === model);
};
