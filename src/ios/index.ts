import devices from '../data/ios.json';
import { IosDevice as Device } from '../types';

function deviceList(): Device[] {
  return devices;
}

/**
 * Get all possible ios devices for a given device identifier.
 */
export const getIosDevices = (identifier: string): Device[] => {
  return deviceList().filter((device) => device.identifier === identifier);
};

export const getIosDevicesFromMarketingName = (
  marketingName: string,
): Device[] => {
  return deviceList().filter((device) => device.generation === marketingName);
};
