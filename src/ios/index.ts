import devices from '../data/ios.json';
import { IosDevice as Device } from '../types';

function deviceList(): Device[] {
  return devices;
}

/**
 * Get the first matching device's generation (aka marketing name) from the list of devices.
 * If no device is found, returns the given model name.
 *
 * @param identifier
 */
export const getIosMarketingName = (identifier: string): string => {
  const device = deviceList().find(
    (device) => device.identifier === identifier,
  );

  let marketingName = device?.generation;

  if (!marketingName) {
    marketingName = identifier;
  }

  return marketingName;
};
