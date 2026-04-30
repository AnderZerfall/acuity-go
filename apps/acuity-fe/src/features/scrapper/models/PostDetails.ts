export interface PostDetails {
  userId: string;
  username: string;
  text: string;
  media: string;
  views: number;
  postDate: Date | null;
}

export const PostDetails = {
  create,
};

function create(init: Partial<PostDetails> = {}): PostDetails {
  return {
    ...init,
    userId: init.userId ?? '',
    username: init.username ?? '',
    text: init.text ?? '',
    media: init.media ?? '',
    views: init.views ?? 0,
    postDate: init.postDate ?? null,
  };
}
