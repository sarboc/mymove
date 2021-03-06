/* This matches the width at which submit buttons go full-width for mobile devices */
export const mobileSize = 481;
export const isProduction = process.env.NODE_ENV === 'production';
export const isDevelopment = process.env.NODE_ENV === 'development';
export const isTest = process.env.NODE_ENV === 'test';
export const NULL_UUID = '00000000-0000-0000-0000-000000000000';
export const ppmInfoPacket = '/downloads/ppm_info_sheet.pdf';

export const hostname = window && window.location && window.location.hostname;
export const isOfficeSite = hostname.startsWith('office') || '';
export const isTspSite = hostname.startsWith('tsp') || '';

export const titleCase = str => {
  return str.charAt(0).toUpperCase() + str.slice(1);
};
