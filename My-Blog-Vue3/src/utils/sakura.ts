/**
 * 樱花飘落效果
 * Sakura falling effect
 */

type RandomOption = "x" | "y" | "s" | "r" | "fnx" | "fny" | "fnr";
type FnX = (x: number, y: number) => number;
type FnY = (x: number, y: number) => number;
type FnR = (r: number) => number;

interface SakuraFunctions {
  x: FnX;
  y: FnY;
  r: FnR;
}

interface SakuraItem {
  x: number;
  y: number;
  s: number;
  r: number;
  fn: SakuraFunctions;
}

let stop: (() => void) | undefined;

const img = new Image();
img.src =
  "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAADIAAAAyCAYAAAAeP4ixAAAACXBIWXMAAAsTAAALEwEAmpwYAAAEsUlEQVR4nO2aTWhcZRTHf2cmSSdJk6ZJmyZNGxqjYl1Y3Ch+ICqCG3EjiIoL0YW6sCAouhQRFVxYEBc+gLhQBBdupVhBLKgoLlz4gVJBBMVKP5I2SZu0zaRJZuY8cl7mDm/uvO/ed2deXAifhWHu3HvP/3/O/5xz58WokC8NN9Pu5kLSdYaGelqAduCbakqJaokqgaqqoP2HLVpQGwIsaLTXiqaqKCQ6V1REbE1BQWOBX+KGF7pAjvUXxbgpZCvSV1TiFJ8fSkTqFBsP3fB0/X+fGhYNe6bLH7J1nqBqDFW5IhYZoSqhGUGtaRWwKkO1hKSc8uKaiGheqaoyNNBQaIJaJaiVYe5fGkqKaFRoMNFAAwWkxFpNx8WVNE7JF5sQa1MUMVbxW5bHsWJc9oVUdYTG4g0EBRSSoDuCQqKhB4hJ3fDIi/UUkqBGlJqoaKNBBYW0pqAr2c+BYqgORXwRPjlqRNVWkUwt/JBMxEHFaE5KQhJpJUFVh8opJBVyqIJ2TjVtI/w5IagYV8R0VNVYVKhNwxrCYgKyRzEOMUn1CKoWklQrJzShYgJqNBRqaKChBqkIankoNlHRhhJrJKSyJqCkdK5IQsAFqmosqhFJpaL0pNDVSqpQXEpXJAGXVW1NNYakKqHLKbYpqI2iqkJXJLq2IkElZYZaM0p1AqoHYT2tVySpdhGhNgmxqEZS+qPEUhw1yOqKuBomhyIaqhoqEpKuq5KQNDoXVBEX0GmgqkJyS/cFNCHVJVHRFXG/FpRSB1EPSTqKJiCqQxqIitg1SE5qGh4vpCqpvhOpGNdQVIeq7FpBUq2YrCiKRhVjPKQqoDqeQkKDhRYECSkBRaVUVlTERYLKQDFJsU5AJpGwGmfT8Ih3RYqp1bRcJ5GQTIyD4lIsI8y/qhvqFdIakoT/C0jEIj7+H/3DpMTaOKoTKCaIakRSKyZpaFxFXFZS0BVJqFGDQlNOeaQyVFVFHOqMUJGEWsRYE1Id1BoK2lCRQnoDSahqCOoN6f4FlVQD3aD0dCUhEZoVqOjfFfMvXBHnDg1FexHZlaqbL7oikpDbNSipFtFJFUmqJUqq1ot6EhLqxLgMDYpxCQlJRExTSGJdLaVEcTQh4AIaEYdaJyQBnQ5qg0G9qIIGVNTkQhIKaKcqYhUhyRrFBNSGVNNEU6mC2nRFTCglNh4aNJ4qQq2oJkNtItIVca3QFUlAHKpxE1DdD00RdTXW1AoKCnGqJk51Qn08XBE1Aar7YW/S85wq4hLiVCfU8UTVqgS1aaTtIqZ2hTZRV6SVAmpNOVXLqE0JSbVh0k1JqsNJSUgaqqCgGk+1TaU6ga5WcqoTGu8KdQqphkKdEKsuoCcpoaZiXA4lxSQwLgppDEki2mlAqYaaoNYU0gQKaqdSSC/qJlwRqxpJtEZS/b9qJLViEhqPG1FDStXQoGos1fFQ0VCSQqxJxQ0PCqq6JqLaNBJCE0GNSVBTQRVVxBqKddSJiiYUdE1A6gRkkgakJqK4lOog1pqAqhhXS0RSQoLqaVIqJuBTSzXVRpKm/gccuE9v+ZJN7gAAAABJRU5ErkJggg==";

// 随机数生成函数
const getRandom = (option: RandomOption): number | SakuraFunctions => {
  let ret: number | SakuraFunctions;
  const random = Math.random();

  switch (option) {
    case "x":
      ret = Math.ceil(random * window.innerWidth);
      break;
    case "y":
      ret = Math.ceil(random * window.innerHeight);
      break;
    case "s":
      ret = random;
      break;
    case "r":
      ret = Math.random() * 6;
      break;
    case "fnx":
    case "fny":
    case "fnr":
      ret = getRandomFn();
      break;
    default:
      ret = random;
  }
  return ret;
};

// 随机函数生成器
const getRandomFn = (): SakuraFunctions => {
  const random = Math.random();
  let fnX: FnX;
  let fnY: FnY;
  let fnR: FnR;

  if (random > 0.6) {
    fnX = (x: number, y: number): number => {
      return x + 0.5 - 1 * (y / window.innerHeight);
    };
  } else if (random > 0.3) {
    fnX = (x: number, y: number): number => {
      return x + 0.5;
    };
  } else {
    fnX = (x: number): number => {
      return x + 0.5 + 1 * Math.random();
    };
  }

  fnY = (_x: number, y: number): number => {
    return y + 1 + 2 * Math.random();
  };

  fnR = (r: number): number => {
    return r + 0.03 * Math.random();
  };

  return {
    x: fnX,
    y: fnY,
    r: fnR,
  };
};

// 创建樱花对象
const createSakura = (): SakuraItem => {
  const fn = getRandomFn();
  return {
    x: getRandom("x") as number,
    y: getRandom("y") as number,
    s: getRandom("s") as number,
    r: getRandom("r") as number,
    fn: fn,
  };
};

// 绘制樱花
const drawSakura = (cxt: CanvasRenderingContext2D, sakura: SakuraItem): void => {
  cxt.save();
  const xc = 40 * sakura.s;
  const yc = 40 * sakura.s;
  cxt.translate(sakura.x, sakura.y);
  cxt.rotate(sakura.r);
  cxt.drawImage(img, 0, 0, xc > 0 ? xc : 0.1, yc > 0 ? yc : 0.1);
  cxt.restore();
};

// 更新樱花位置
const updateSakura = (sakura: SakuraItem): SakuraItem => {
  sakura.x = sakura.fn.x(sakura.x, sakura.y);
  sakura.y = sakura.fn.y(sakura.x, sakura.y);
  sakura.r = sakura.fn.r(sakura.r);

  if (
    sakura.x > window.innerWidth ||
    sakura.x < 0 ||
    sakura.y > window.innerHeight ||
    sakura.y < 0
  ) {
    sakura.x = getRandom("x") as number;
    sakura.y = 0;
    sakura.s = getRandom("s") as number;
    sakura.r = getRandom("r") as number;
  }

  return sakura;
};

// 动画循环
const sakuraAnimate = (
  cxt: CanvasRenderingContext2D,
  canvas: HTMLCanvasElement,
  sakuraList: SakuraItem[]
): void => {
  let requestId: number;

  const loop = (): void => {
    cxt.clearRect(0, 0, canvas.width, canvas.height);
    sakuraList.forEach((sakura, index) => {
      sakura = updateSakura(sakura);
      sakuraList[index] = sakura;
      drawSakura(cxt, sakura);
    });
    requestId = requestAnimationFrame(loop);
  };

  loop();

  stop = (): void => {
    cancelAnimationFrame(requestId);
  };
};

// 初始化
const init = (): void => {
  // 创建画布
  const canvas = document.createElement("canvas");
  canvas.id = "sakura-canvas";
  canvas.style.cssText = `
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    pointer-events: none;
    z-index: 9999;
  `;
  document.body.appendChild(canvas);

  canvas.width = window.innerWidth;
  canvas.height = window.innerHeight;

  const cxt = canvas.getContext("2d");
  if (!cxt) return;

  // 创建樱花数组
  const sakuraList: SakuraItem[] = [];
  const sakuraNum = 50;
  for (let i = 0; i < sakuraNum; i++) {
    sakuraList.push(createSakura());
  }

  // 开始动画
  sakuraAnimate(cxt, canvas, sakuraList);

  // 响应窗口大小变化
  window.addEventListener("resize", () => {
    canvas.width = window.innerWidth;
    canvas.height = window.innerHeight;
  });
};

// 等待图片加载完成后启动
img.onload = (): void => {
  init();
};

export { stop };
