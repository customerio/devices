import { getMarketingName } from './index';
import { OsType } from '../types';

describe('common', () => {
  test('ios: should get the marketing name', () => {
    const marketingName = getMarketingName('iPhone14,2', OsType.IOS);
    expect(marketingName).toBe('iPhone 13 Pro');
  });

  test('android: should get the marketing name', () => {
    const marketingName = getMarketingName('SM-A300H', OsType.ANDROID);
    expect(marketingName).toBe('Galaxy A3');
  });

  test('should return the model name if no marketing name is found', () => {
    let marketingName = getMarketingName('iUNKNOWN', OsType.IOS);
    expect(marketingName).toBe('iUNKNOWN');
    marketingName = getMarketingName('SM-UNKNOWN', OsType.ANDROID);
    expect(marketingName).toBe('SM-UNKNOWN');
  });
});
