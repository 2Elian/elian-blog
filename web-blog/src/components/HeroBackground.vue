<template>
  <div class="space-bg" ref="containerRef">
    <canvas ref="canvasRef" class="particle-canvas" />
    <div class="space-layer">
      <div class="nebula nebula-1"></div>
      <div class="nebula nebula-2"></div>
      <div class="nebula nebula-3"></div>
      <div class="planet jupiter"></div>
      <div class="planet earth"></div>
      <div class="planet moon"></div>
      <div class="star-dot dot-1"></div>
      <div class="star-dot dot-2"></div>
      <div class="star-dot dot-3"></div>
      <div class="star-trail"></div>
      <div class="star-trail star-trail-2"></div>
      <div class="cursor-glow" ref="glowRef"></div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue'

const containerRef = ref<HTMLDivElement>()
const canvasRef = ref<HTMLCanvasElement>()
const glowRef = ref<HTMLDivElement>()

let animId = 0
let particles: Particle[] = []
let mouse = { x: -1000, y: -1000 }

const PARTICLE_COUNT = 200
const CONNECTION_DISTANCE = 100
const PARTICLE_SPEED = 0.4

class Particle {
  x: number
  y: number
  vx: number
  vy: number
  radius: number
  color: string
  opacity: number
  twinkleSpeed: number
  twinkleOffset: number
  w: number
  h: number

  constructor(w: number, h: number) {
    this.w = w
    this.h = h
    this.x = Math.random() * w
    this.y = Math.random() * h
    this.vx = (Math.random() - 0.5) * PARTICLE_SPEED
    this.vy = (Math.random() - 0.5) * PARTICLE_SPEED
    this.radius = Math.random() * 1.8 + 0.3
    const colors = [
      '200, 220, 255', '225, 210, 255', '255, 240, 220',
      '180, 200, 230', '210, 190, 220', '190, 210, 240'
    ]
    this.color = colors[Math.floor(Math.random() * colors.length)]
    this.opacity = Math.random() * 0.8 + 0.2
    this.twinkleSpeed = Math.random() * 0.02 + 0.005
    this.twinkleOffset = Math.random() * Math.PI * 2
  }

  update() {
    this.x += this.vx
    this.y += this.vy
    const dx = this.x - mouse.x
    const dy = this.y - mouse.y
    const dist = Math.sqrt(dx * dx + dy * dy)
    if (dist < 130) {
      const angle = Math.atan2(dy, dx)
      const force = (130 - dist) / 130
      this.vx += Math.cos(angle) * force * 0.15
      this.vy += Math.sin(angle) * force * 0.15
    }
    const speed = Math.sqrt(this.vx * this.vx + this.vy * this.vy)
    if (speed > PARTICLE_SPEED * 1.8) {
      this.vx = (this.vx / speed) * PARTICLE_SPEED * 1.8
      this.vy = (this.vy / speed) * PARTICLE_SPEED * 1.8
    }
    if (this.x < -30) this.x = this.w + 30
    if (this.x > this.w + 30) this.x = -30
    if (this.y < -30) this.y = this.h + 30
    if (this.y > this.h + 30) this.y = -30
  }

  draw(ctx: CanvasRenderingContext2D, time: number) {
    const twinkle = Math.sin(time * this.twinkleSpeed + this.twinkleOffset) * 0.3 + 0.7
    const currentOpacity = this.opacity * twinkle
    ctx.beginPath()
    ctx.arc(this.x, this.y, this.radius, 0, Math.PI * 2)
    ctx.fillStyle = `rgba(${this.color}, ${currentOpacity})`
    ctx.fill()
    if (this.radius > 1.3) {
      ctx.beginPath()
      ctx.arc(this.x, this.y, this.radius * 2.2, 0, Math.PI * 2)
      ctx.fillStyle = `rgba(${this.color}, ${currentOpacity * 0.12})`
      ctx.fill()
    }
  }
}

function onMouseMove(e: MouseEvent) {
  mouse.x = e.clientX
  mouse.y = e.clientY
  if (glowRef.value) {
    glowRef.value.style.left = e.clientX + 'px'
    glowRef.value.style.top = e.clientY + 'px'
    glowRef.value.style.opacity = '1'
  }
  const moveX = (e.clientX - window.innerWidth / 2) / window.innerWidth
  const moveY = (e.clientY - window.innerHeight / 2) / window.innerHeight
  if (containerRef.value) {
    const nebulae = containerRef.value.querySelectorAll('.nebula')
    const planets = containerRef.value.querySelectorAll('.planet')
    const dots = containerRef.value.querySelectorAll('.star-dot')
    nebulae.forEach((el, i) => { (el as HTMLElement).style.transform = `translate(${moveX * (12 + i * 3)}px, ${moveY * (12 + i * 3)}px)` })
    planets.forEach((el, i) => { (el as HTMLElement).style.transform = `translate(${moveX * (20 + i * 6)}px, ${moveY * (20 + i * 6)}px)` })
    dots.forEach((el, i) => { (el as HTMLElement).style.transform = `translate(${moveX * (10 + i * 2)}px, ${moveY * (10 + i * 2)}px)` })
  }
}

function onMouseLeave() {
  mouse.x = -1000
  mouse.y = -1000
  if (glowRef.value) glowRef.value.style.opacity = '0'
  if (containerRef.value) {
    containerRef.value.querySelectorAll('.nebula, .planet, .star-dot').forEach(el => {
      (el as HTMLElement).style.transform = 'translate(0px, 0px)'
    })
  }
}

onMounted(() => {
  const canvas = canvasRef.value!
  const ctx = canvas.getContext('2d')!
  let w = 0, h = 0

  function resize() {
    w = window.innerWidth
    h = window.innerHeight
    canvas.width = w
    canvas.height = h
    particles.forEach(p => { p.w = w; p.h = h })
  }
  resize()
  window.addEventListener('resize', resize)
  document.addEventListener('mousemove', onMouseMove)
  document.addEventListener('mouseleave', onMouseLeave)

  particles = Array.from({ length: PARTICLE_COUNT }, () => new Particle(w, h))

  function drawConnections() {
    for (let i = 0; i < particles.length; i++) {
      for (let j = i + 1; j < particles.length; j++) {
        const dx = particles[i].x - particles[j].x
        const dy = particles[i].y - particles[j].y
        const dist = Math.sqrt(dx * dx + dy * dy)
        if (dist < CONNECTION_DISTANCE) {
          const opacity = 1 - dist / CONNECTION_DISTANCE
          ctx.beginPath()
          ctx.moveTo(particles[i].x, particles[i].y)
          ctx.lineTo(particles[j].x, particles[j].y)
          ctx.strokeStyle = `rgba(255, 255, 255, ${opacity * 0.07})`
          ctx.lineWidth = 0.4
          ctx.stroke()
        }
      }
    }
  }

  function animate(time: number) {
    ctx.clearRect(0, 0, w, h)
    for (const p of particles) {
      p.update()
      p.draw(ctx, time)
    }
    drawConnections()
    animId = requestAnimationFrame(animate)
  }
  animId = requestAnimationFrame(animate)
})

onUnmounted(() => {
  cancelAnimationFrame(animId)
  document.removeEventListener('mousemove', onMouseMove)
  document.removeEventListener('mouseleave', onMouseLeave)
})
</script>

<style scoped lang="scss">
.space-bg {
  position: absolute;
  inset: 0;
  overflow: hidden;
}

.particle-canvas {
  position: absolute;
  inset: 0;
  width: 100%;
  height: 100%;
  z-index: 0;
}

.space-layer {
  position: absolute;
  inset: 0;
  z-index: 1;
  pointer-events: none;
}

// Nebulae
.nebula {
  position: absolute;
  border-radius: 50%;
  filter: blur(100px);
  opacity: 0.45;
  will-change: transform;
  animation: floatNebula 18s infinite alternate ease-in-out;
  mix-blend-mode: screen;
}
.nebula-1 {
  width: 70vw; height: 70vw; max-width: 900px; max-height: 900px;
  background: radial-gradient(circle, #1e3a5f 0%, transparent 70%);
  top: -15%; left: -10%;
  animation-duration: 20s;
}
.nebula-2 {
  width: 80vw; height: 80vw; max-width: 1000px; max-height: 1000px;
  background: radial-gradient(circle, #2a1a4a 0%, transparent 70%);
  bottom: -20%; right: -12%;
  animation-duration: 22s; animation-delay: -6s;
}
.nebula-3 {
  width: 60vw; height: 60vw; max-width: 700px; max-height: 700px;
  background: radial-gradient(circle, #0f2b3d 0%, transparent 70%);
  top: 45%; left: 40%; opacity: 0.35;
  animation-duration: 25s; animation-delay: -10s;
}

@keyframes floatNebula {
  0% { transform: translate(0, 0); }
  100% { transform: translate(30px, -20px); }
}

// Planets
.planet {
  position: absolute;
  border-radius: 50%;
  will-change: transform;
  animation: floatPlanet 24s infinite alternate ease-in-out;
  box-shadow: 0 0 30px rgba(100, 150, 255, 0.3), inset 0 -10px 20px rgba(0,0,0,0.5);
}
.jupiter {
  width: 130px; height: 130px; top: 15%; right: 20%;
  background: radial-gradient(circle at 30% 30%, #d4a373, #8b5a2b 50%, #4a2e15 80%),
  repeating-linear-gradient(0deg, rgba(210, 160, 100, 0.4) 0px, rgba(210, 160, 100, 0.4) 6px,
  rgba(180, 120, 60, 0.3) 6px, rgba(180, 120, 60, 0.3) 12px);
  animation-duration: 28s; animation-delay: -3s;
}
.earth {
  width: 85px; height: 85px; bottom: 20%; left: 15%;
  background: radial-gradient(circle at 40% 40%, #4da6ff, #0066cc 60%, #001a33 90%),
  radial-gradient(circle at 60% 30%, #66cc66 8%, transparent 12%),
  radial-gradient(circle at 25% 70%, #66cc66 6%, transparent 10%);
  animation-duration: 22s; animation-delay: -8s;
}
.moon {
  width: 60px; height: 60px; top: 55%; left: 28%;
  background: radial-gradient(circle at 35% 35%, #e0e0e0, #888888 70%, #444444 100%);
  animation-duration: 20s; animation-delay: -12s;
  box-shadow: 0 0 25px rgba(200,200,200,0.4), inset 0 -6px 12px rgba(0,0,0,0.6);
}
.planet::after {
  content: '';
  position: absolute; top: 50%; left: 50%;
  transform: translate(-50%, -50%) rotateX(75deg) rotateY(20deg);
  border-radius: 50%;
  border: 2px solid rgba(255, 255, 255, 0.25);
  box-shadow: 0 0 15px rgba(255, 255, 255, 0.2);
  animation: ringPulse 8s infinite alternate ease-in-out;
  pointer-events: none;
}
.jupiter::after { width: 170px; height: 170px; border-width: 3px; border-color: rgba(210, 180, 140, 0.4); }
.earth::after { width: 115px; height: 115px; border-width: 1.5px; border-color: rgba(100, 180, 255, 0.5); }
.moon::after { width: 85px; height: 85px; border-width: 1px; border-style: dashed; border-color: rgba(220, 220, 255, 0.5); animation-duration: 5s; }

@keyframes floatPlanet {
  0% { transform: translate(0, 0) scale(1); }
  33% { transform: translate(25px, -45px) scale(1.03); }
  66% { transform: translate(-20px, 30px) scale(0.97); }
  100% { transform: translate(15px, -25px) scale(1.01); }
}
@keyframes ringPulse {
  0% { opacity: 0.5; transform: translate(-50%, -50%) rotateX(75deg) rotateY(20deg) scale(1); }
  100% { opacity: 0.9; transform: translate(-50%, -50%) rotateX(75deg) rotateY(20deg) scale(1.02); }
}

// Star dots
.star-dot {
  position: absolute; border-radius: 50%; background: #b0e0ff;
  box-shadow: 0 0 12px #7bb3d9, 0 0 25px rgba(100, 150, 255, 0.5);
  animation: twinkle 10s infinite alternate ease-in-out;
}
.dot-1 { width: 8px; height: 8px; top: 28%; left: 18%; background: #c4d2ff; animation-duration: 8s; }
.dot-2 { width: 5px; height: 5px; top: 72%; right: 18%; background: #ffe0b0; animation-duration: 12s; animation-delay: -3s; }
.dot-3 { width: 6px; height: 6px; bottom: 25%; left: 55%; background: #c9e4ff; animation-duration: 14s; animation-delay: -5s; }
@keyframes twinkle {
  0% { opacity: 0.6; transform: scale(1); }
  100% { opacity: 1; transform: scale(1.6); }
}

// Star trails
.star-trail {
  position: absolute; bottom: 0; left: 0; width: 100%; height: 2px;
  background: linear-gradient(90deg, transparent, #4a6fa5, #b0c4de, transparent);
  animation: slideLine 8s linear infinite; filter: blur(2px); opacity: 0.5;
}
.star-trail-2 { top: 0; bottom: auto; animation-duration: 12s; background: linear-gradient(90deg, transparent, #5c4a7a, #9b8ec4, transparent); animation-delay: -4s; }
@keyframes slideLine {
  0% { transform: translateX(-105%); }
  100% { transform: translateX(105%); }
}

// Cursor glow
.cursor-glow {
  position: absolute; width: 450px; height: 450px; border-radius: 50%;
  background: radial-gradient(circle, rgba(173, 200, 255, 0.12) 0%, transparent 70%);
  filter: blur(70px); pointer-events: none; z-index: 2;
  transform: translate(-50%, -50%); transition: opacity 0.3s ease; opacity: 0;
}

@media (max-width: 768px) {
  .jupiter { width: 90px; height: 90px; }
  .earth { width: 65px; height: 65px; }
  .moon { width: 45px; height: 45px; }
  .cursor-glow { width: 200px; height: 200px; }
}
</style>
