import { REDDIT_REGEX, TWITTER_REGEX } from '../../form/constants/url-patterns';
import { SupportedSocial } from '../constants/SupportedSocials';

export const SocialUrlMapper = {
  toSocials,
};

function toSocials(url: string): SupportedSocial | null {
  if (REDDIT_REGEX.test(url)) {
    return SupportedSocial.Reddit;
  }

  if (TWITTER_REGEX.test(url)) {
    return SupportedSocial.X;
  }

  return null;
}
