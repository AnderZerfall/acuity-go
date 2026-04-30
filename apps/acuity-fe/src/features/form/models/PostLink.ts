import { z } from 'zod';
import { REDDIT_REGEX, TWITTER_REGEX } from '../constants/url-patterns';

export interface PostLink {
  url: string;
}

export const PostLink = {
  create,
  schema: z.object({
    url: z
      .string()
      .trim()
      .refine((val) => TWITTER_REGEX.test(val) || REDDIT_REGEX.test(val), {
        message: 'URL has to be from Twitter or Reddit post',
      }),
  }),
};

function create(init: Partial<PostLink> = {}): PostLink {
  return { ...init, url: init?.url ?? '' };
}
