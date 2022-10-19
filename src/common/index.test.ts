import { getMarketingName } from './index';
import { Platform } from '../types';

describe('common', () => {
  test('ios: should get the marketing name', () => {
    let marketingName = getMarketingName('iPhone14,2', Platform.IOS);
    expect(marketingName).toBe('iPhone 13 Pro');

    // case insensitive
    marketingName = getMarketingName('iphone14,3', Platform.IOS);
    expect(marketingName).toBe('iPhone 13 Pro Max');
  });

  test('android: should get the marketing name', () => {
    let marketingName = getMarketingName('SM-A300H', Platform.ANDROID);
    expect(marketingName).toBe('Galaxy A3');

    // case insensitive
    marketingName = getMarketingName('2201123g', Platform.ANDROID);
    expect(marketingName).toBe('Xiaomi 12');
  });

  test('should return the model name if no marketing name is found', () => {
    let marketingName = getMarketingName('iUNKNOWN', Platform.IOS);
    expect(marketingName).toBe('iUNKNOWN');
    marketingName = getMarketingName('SM-UNKNOWN', Platform.ANDROID);
    expect(marketingName).toBe('SM-UNKNOWN');
  });
});
