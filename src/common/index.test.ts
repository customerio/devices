import { getMarketingName } from './index';

describe('common', () => {
  test('ios: should get the marketing name', () => {
    let marketingName = getMarketingName('iPhone14,2');
    expect(marketingName).toBe('iPhone 13 Pro');

    // case insensitive
    marketingName = getMarketingName('iphone14,3');
    expect(marketingName).toBe('iPhone 13 Pro Max');

    marketingName = getMarketingName('iPhone19,2');
    expect(marketingName).toBe('iPhone 18 Pro');

    marketingName = getMarketingName('iPhone19,3');
    expect(marketingName).toBe('iPhone 18 Pro Max');

    marketingName = getMarketingName('iPhone19,7');
    expect(marketingName).toBe('iPhone 18 Pro Max');

    marketingName = getMarketingName('iPhone19,4');
    expect(marketingName).toBe('iPhone Duo');

    marketingName = getMarketingName('iPhone17,5');
    expect(marketingName).toBe('iPhone 16e');

    marketingName = getMarketingName('iPhone18,1');
    expect(marketingName).toBe('iPhone 17 Pro');

    marketingName = getMarketingName('iPhone18,4');
    expect(marketingName).toBe('iPhone Air');

    marketingName = getMarketingName('Watch7,17');
    expect(marketingName).toBe('Apple Watch Series 11');

    marketingName = getMarketingName('Watch7,12');
    expect(marketingName).toBe('Apple Watch Ultra 3');
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
