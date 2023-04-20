import type { Price, CurrencyName } from '#/gen/models';

export const displayPrice = (price: Price | undefined, currencyName: CurrencyName | undefined): string => {
  if (typeof price === 'undefined' || typeof currencyName === 'undefined') {
    return '';
  }
  return `${price.toLocaleString()} ${currencyName}`;
}