import { ScrapperService } from '../features/scrapper/service/ScrapperService';

console.log('content-worker started');

// Signal ready via one-time message
chrome.runtime.sendMessage({ type: 'CONTENT_READY' });

// Listen for port connection from service worker
chrome.runtime.onConnect.addListener((port) => {
  if (port.name !== 'SCRAPE_PORT') return;

  port.onMessage.addListener(async (message) => {
    if (message.type !== 'SCRAPE_TAB') return;

    try {
      await waitForPageIdle();
      const data = ScrapperService.scrape(
        message.target,
        document.documentElement,
      );
      port.postMessage({ success: true, data });
    } catch (err) {
      console.error('Scrape error:', err);
      port.postMessage({
        success: false,
        error: err instanceof Error ? err.message : 'Unknown error',
      });
    }
  });
});

function waitForPageIdle(timeout = 15000): Promise<void> {
  return new Promise((resolve) => {
    if (document.readyState === 'complete') {
      setTimeout(resolve, 2000);
      return;
    }

    window.addEventListener(
      'load',
      () => {
        setTimeout(resolve, 2000);
      },
      { once: true },
    );

    setTimeout(resolve, timeout);
  });
}
