import type { Date } from '#/gen/models';

export const displayDate = (date: Date | undefined): string => {
  if (typeof date !== 'undefined') {
    if (typeof date.year === 'number' && typeof date.month === 'number' && typeof date.day === 'number') {
      const dateObject = new Date(`${date.year}-${date.month}-${date.day}`);
      const ret = dateObject.toLocaleDateString();
      return ret;
    }
  }
  return '';
};