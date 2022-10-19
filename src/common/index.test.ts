import { getMarketingName } from './index';

describe('common', () => {
  test('ios: should get the marketing name', () => {
    let marketingName = getMarketingName('iPhone14,2');
    expect(marketingName).toBe('iPhone 13 Pro');

    // case insensitive
    marketingName = getMarketingName('iphone14,3');
    expect(marketingName).toBe('iPhone 13 Pro Max');
  });

  test('android: should get the marketing name', () => {
    let marketingName = getMarketingName('SM-A300H');
    expect(marketingName).toBe('Galaxy A3');

    // case insensitive
    marketingName = getMarketingName('2201123g');
    expect(marketingName).toBe('Xiaomi 12');
  });

  test('should return null if no marketing name is found', () => {
    const marketingName = getMarketingName('iUNKNOWN');
    expect(marketingName).toBeNull();
  });
});
