import devices from '../data/ios.json';
import { IosDevice as Device } from '../types';

function deviceList(): Device[] {
  return devices;
}

/**
 * Get all possible ios devices for a given device identifier.
 */
export const getIosDevices = (identifier: string): Device[] => {
  return deviceList().filter(
    (device) => device.identifier.toLowerCase() === identifier.toLowerCase(),
  );
};

export const getIosDevicesFromMarketingName = (
  marketingName: string,
): Device[] => {
  return deviceList().filter(
    (device) => device.generation.toLowerCase() === marketingName.toLowerCase(),
  );
};
