interface Particle {
  x: number;
  y: number;
  color: string;
  radius: number;
  endPos: { x: number; y: number };
  draw: () => void;
}

interface Circle {
  x: number;
  y: number;
  color: string;
  radius: number;
  alpha: number;
  lineWidth: number;
  draw: () => void;
}

interface AnimeAnimatable {
  target: Particle | Circle;
}

interface LocalAnimeCallbackParam {
  animatables: AnimeAnimatable[];
}

function debounce<T extends (...args: unknown[]) => void>(
  fn: T,
  delay: number
): (...args: Parameters<T>) => void {
  let timer: ReturnType<typeof setTimeout> | undefined;
  return function (this: unknown, ...args: Parameters<T>): void {
    clearTimeout(timer);
    timer = setTimeout(() => {
      fn.apply(this, args);
    }, delay);
  };
}

export default function (): void {
  // 动画效果
  const canvasEl = document.querySelector("#mousedown") as HTMLCanvasElement | null;
  if (!canvasEl) return;

  const ctx = canvasEl.getContext("2d", { willReadFrequently: true });
  if (!ctx) return;

  // 使用局部常量来保证TypeScript知道它们非null
  const canvas = canvasEl;
  const context = ctx;

  const numberOfParticles = 30;
  let pointerX = 0;
  let pointerY = 0;
  const tap = "mousedown";
  const colors = ["#FF1461", "#18FF92", "#5A87FF", "#FBF38C"];

  const setCanvasSize = debounce(function (): void {
    canvas.width = 2 * window.innerWidth;
    canvas.height = 2 * window.innerHeight;
    canvas.style.width = window.innerWidth + "px";
    canvas.style.height = window.innerHeight + "px";
    const newContext = canvas.getContext("2d", { willReadFrequently: true });
    if (newContext) {
      newContext.scale(2, 2);
    }
  }, 500);

  const render = window.anime({
    duration: 1 / 0,
    update: function (): void {
      context.clearRect(0, 0, canvas.width, canvas.height);
    },
  });

  function updateCoords(e: MouseEvent | TouchEvent): void {
    if ("touches" in e) {
      pointerX = e.touches[0].clientX - canvas.getBoundingClientRect().left;
      pointerY = e.touches[0].clientY - canvas.getBoundingClientRect().top;
    } else {
      pointerX = e.clientX - canvas.getBoundingClientRect().left;
      pointerY = e.clientY;
    }
  }

  function setParticleDirection(e: { x: number; y: number }): { x: number; y: number } {
    const t = (window.anime.random(0, 360) * Math.PI) / 180;
    const a = window.anime.random(50, 180);
    const n = [-1, 1][window.anime.random(0, 1)] * a;
    return {
      x: e.x + n * Math.cos(t),
      y: e.y + n * Math.sin(t),
    };
  }

  function createParticle(e: number, t: number): Particle {
    const a: Particle = {
      x: e,
      y: t,
      color: colors[window.anime.random(0, colors.length - 1)],
      radius: window.anime.random(16, 32),
      endPos: { x: 0, y: 0 },
      draw: function (): void {
        context.beginPath();
        context.arc(a.x, a.y, a.radius, 0, 2 * Math.PI, true);
        context.fillStyle = a.color;
        context.fill();
      },
    };
    a.endPos = setParticleDirection(a);
    return a;
  }

  function createCircle(e: number, t: number): Circle {
    const a: Circle = {
      x: e,
      y: t,
      color: "#F00",
      radius: 0.1,
      alpha: 0.5,
      lineWidth: 6,
      draw: function (): void {
        context.globalAlpha = a.alpha;
        context.beginPath();
        context.arc(a.x, a.y, a.radius, 0, 2 * Math.PI, true);
        context.lineWidth = a.lineWidth;
        context.strokeStyle = a.color;
        context.stroke();
        context.globalAlpha = 1;
      },
    };
    return a;
  }

  function renderParticle(anim: unknown): void {
    const e = anim as LocalAnimeCallbackParam;
    for (let t = 0; t < e.animatables.length; t++) {
      e.animatables[t].target.draw();
    }
  }

  function animateParticles(e: number, t: number): void {
    const a = createCircle(e, t);
    const n: Particle[] = [];
    for (let i = 0; i < numberOfParticles; i++) {
      n.push(createParticle(e, t));
    }
    window.anime
      .timeline()
      .add({
        targets: n,
        x: function (particle: Particle): number {
          return particle.endPos.x;
        },
        y: function (particle: Particle): number {
          return particle.endPos.y;
        },
        radius: 0.1,
        duration: window.anime.random(1200, 1800),
        easing: "easeOutExpo",
        update: renderParticle,
      })
      .add({
        targets: a,
        radius: window.anime.random(80, 160),
        lineWidth: 0,
        alpha: {
          value: 0,
          easing: "linear",
          duration: window.anime.random(600, 800),
        },
        duration: window.anime.random(1200, 1800),
        easing: "easeOutExpo",
        update: renderParticle,
        offset: 0,
      });
  }

  document.addEventListener(
    tap,
    function (e: MouseEvent): void {
      const target = e.target as HTMLElement;
      if (
        target.id !== "sidebar" &&
        target.id !== "toggle-sidebar" &&
        target.nodeName !== "A" &&
        target.nodeName !== "IMG"
      ) {
        render.play();
        updateCoords(e);
        animateParticles(pointerX, pointerY);
      }
    },
    false
  );
  setCanvasSize();
  window.addEventListener("resize", setCanvasSize, false);
}
