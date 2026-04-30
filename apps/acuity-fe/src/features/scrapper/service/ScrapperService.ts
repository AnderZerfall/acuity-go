import { ApplicationError } from '../../../utils/error/ApplicationError';
import { PostDetails } from '../models/PostDetails';

export const ScrapperService = {
  scrape,
};

function scrape(target: string, html: HTMLElement): PostDetails {
  switch (target) {
    case 'reddit':
      return scrapeReddit(html);
    case 'x':
      return scrapeTwitter(html);
    default:
      throw ApplicationError.create({
        name: 'Scrapper',
        message: 'Scrapper does not support this social type yet',
        details: {
          target: [target],
        },
      });
  }
}

// currently supports only photos
function scrapeTwitter(html: HTMLElement): PostDetails {
  const tweet = document.querySelector('[data-testid="tweet"]');

  if (!tweet) {
    throw ApplicationError.create({
      name: 'Scrapper',
      message: 'Scrapper was unable to find the post',
      details: {
        target: ['x', '[data-testid="tweet"]'],
      },
    });
  }

  const tweetDate: HTMLTimeElement | null = tweet.querySelector('time') ?? null;
  const tweetText = tweet
    .querySelector('[data-testid="tweetText"]')
    ?.querySelector('span');
  const tweetUsername = tweet
    .querySelector('[data-testid="User-Name"]')
    ?.querySelector('span');

  const tweetMedia = tweet.querySelector('[data-testid="tweetMedia"]');
  const photo: HTMLImageElement | null = tweetMedia
    ? tweetMedia.querySelector('img')
    : null;

  return PostDetails.create({
    userId: tweetUsername?.innerText ?? '',
    username: tweetUsername?.innerText ?? '',
    text: tweetText?.innerText ?? '',
    postDate: tweetDate ? new Date(tweetDate.dateTime) : null,
    media: photo?.src ?? '',
  });
}

// currently supports only photos
function scrapeReddit(html: HTMLElement): PostDetails {
  const reddit = document.querySelector('shreddit-post');

  if (!reddit) {
    throw ApplicationError.create({
      name: 'Scrapper',
      message: 'Scrapper was unable to find the post',
      details: {
        target: ['reddit', 'shreddit-post'],
      },
    });
  }

  const redditDate: HTMLTimeElement | null =
    reddit.querySelector('time') ?? null;

  const redditBodyText = reddit.querySelector<HTMLElement>(
    'shreddit-post-text-body',
  )?.innerText;
  const redditTitleText = reddit.querySelector('h1')?.innerText;
  const redditUsername = reddit.getAttribute('author');
  const photo: HTMLImageElement | null =
    reddit.querySelector('img#post-image') ?? null;

  const video: HTMLVideoElement | null = reddit.querySelector('video') ?? null;

  return PostDetails.create({
    userId: redditUsername ?? '',
    username: redditUsername ?? '',
    text: `${redditTitleText} ${redditBodyText}`,
    postDate: redditDate ? new Date(redditDate.dateTime) : null,
    media: photo?.src || video?.src || '',
  });
}
