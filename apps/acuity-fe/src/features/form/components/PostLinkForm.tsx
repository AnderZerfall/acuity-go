import { zodResolver } from '@hookform/resolvers/zod';
import { useState } from 'react';
import { useForm } from 'react-hook-form';
import { Button } from '../../../app/components/button/Button';
import { URLField } from '../../../app/components/input/URLField';
import { SocialUrlMapper } from '../../scrapper/mappers/SocialUrlMapper';
import { PostLink } from '../models/PostLink';

export function PostLinkForm() {
  const {
    control,
    handleSubmit,
    formState: { errors },
  } = useForm<PostLink>({
    resolver: zodResolver(PostLink.schema),
    defaultValues: PostLink.create(),
    reValidateMode: 'onChange',
  });
  const [response, setResponse] = useState('');

  const onSubmit = (data: PostLink) => {
    const target = SocialUrlMapper.toSocials(data.url);

    if (!target) {
      return;
    }

    chrome.runtime.sendMessage(
      { type: 'SCRAPE', url: data.url, target },
      (response) => {
        if (response.success) {
          console.log('FORM: ', response.data);
        }

        console.log('FORM: ', response);

        setResponse(response);
      },
    );
  };

  const handleFormSumit = handleSubmit((data) => onSubmit(data));

  return (
    <form onSubmit={handleFormSumit}>
      <URLField
        name="url"
        control={control}
        label="Pass link to the post"
        placeholder="https://x.com/status/..."
        error={errors.url}
      />

      <Button type="submit">Analyze Post</Button>

      {'Test: ' + JSON.stringify(response, null, 2)}
    </form>
  );
}
