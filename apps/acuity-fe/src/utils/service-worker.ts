chrome.runtime.onMessage.addListener((message, _sender, sendResponse) => {
  if (message.type === 'SCRAPE') {
    (async () => {
      try {
        const data = await openHiddenTabAndScrape(message.url, message.target);
        sendResponse({ success: true, data });
      } catch (err) {
        sendResponse({
          success: false,
          error: err instanceof Error ? err.message : 'Unknown error',
        });
      }
    })();
    return true;
  }
});

async function openHiddenTabAndScrape(
  url: string,
  target: string,
): Promise<any> {
  const tab = await chrome.tabs.create({ url, active: false });
  const tabId = tab.id;
  if (!tabId) throw new Error('Failed to create tab');

  const contentReadyPromise = waitForContentReady(tabId);
  await waitForTabLoad(tabId);

  await chrome.scripting.executeScript({
    target: { tabId },
    files: ['content.js'],
  });

  await contentReadyPromise;

  const data = await scrapeViaPort(tabId, target);
  await chrome.tabs.remove(tabId);
  return data;
}

function scrapeViaPort(tabId: number, target: string): Promise<any> {
  return new Promise((resolve, reject) => {
    const port = chrome.tabs.connect(tabId, { name: 'SCRAPE_PORT' });

    port.onMessage.addListener((response) => {
      port.disconnect();
      if (response.success) {
        resolve(response.data);
      } else {
        reject(new Error(response.error));
      }
    });

    port.onDisconnect.addListener(() => {
      const err = chrome.runtime.lastError;
      reject(new Error(err?.message ?? 'Port disconnected unexpectedly'));
    });

    port.postMessage({ type: 'SCRAPE_TAB', target });
  });
}

function waitForTabLoad(tabId: number): Promise<void> {
  return new Promise((resolve) => {
    const listener = (id: number, changeInfo: chrome.tabs.OnUpdatedInfo) => {
      if (id === tabId && changeInfo.status === 'complete') {
        chrome.tabs.onUpdated.removeListener(listener);
        resolve();
      }
    };
    chrome.tabs.onUpdated.addListener(listener);
  });
}

function waitForContentReady(tabId: number): Promise<void> {
  return new Promise((resolve) => {
    const listener = (message: any, sender: chrome.runtime.MessageSender) => {
      if (message.type === 'CONTENT_READY' && sender.tab?.id === tabId) {
        chrome.runtime.onMessage.removeListener(listener);
        resolve();
      }
    };
    chrome.runtime.onMessage.addListener(listener);
  });
}
