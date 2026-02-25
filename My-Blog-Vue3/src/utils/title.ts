let titleTime: ReturnType<typeof setTimeout> | undefined;
const OriginTitle: string = document.title;

document.addEventListener("visibilitychange", function (): void {
  if (document.hidden) {
    document.title = "w(ﾟДﾟ)w 不要走！再看看嘛！😭";
    if (titleTime) {
      clearTimeout(titleTime);
    }
  } else {
    document.title = "♪(^∇^*)欢迎肥来！🥰";
    titleTime = setTimeout(function (): void {
      document.title = OriginTitle;
    }, 2000);
  }
});
