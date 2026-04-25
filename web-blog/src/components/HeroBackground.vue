<template>
  <div class="space-bg" ref="containerRef">
    <canvas ref="canvasRef" class="particle-canvas" />
    <div class="space-layer">
      <div class="nebula nebula-1"></div>
      <div class="nebula nebula-2"></div>
      <div class="nebula nebula-3"></div>
      <div class="planet jupiter"></div>
      <div class="planet saturn"><div class="saturn-ring"></div></div>
      <div class="planet earth"></div>
      <div class="planet mars"></div>
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
let constellationNodes: number[] = []
let constellationConnections: { i: number; j: number; phase: number; speed: number }[] = []
let explosions: Explosion[] = []
let mouse = { x: -1000, y: -1000, clickX: -1, clickY: -1, clicked: false }
let alive = false

const PARTICLE_COUNT = 230
const CONNECTION_DISTANCE = 110
const PARTICLE_SPEED = 0.45
const CONSTELLATION_NODE_COUNT = 45
const CONSTELLATION_CONNECT_DISTANCE = 190

// Planet click targets (relative to canvas)
const planetSelectors = ['jupiter', 'saturn', 'earth', 'mars', 'moon']

interface Explosion {
  x: number
  y: number
  particles: { x: number; y: number; vx: number; vy: number; life: number; decay: number; color: string; size: number }[]
  shockwaveRadius: number
  shockwaveMaxRadius: number
  shockwaveLife: number
  shockwaveDecay: number
}

class Particle {
  x: number; y: number; vx: number; vy: number; radius: number
  color: string; opacity: number; twinkleSpeed: number; twinkleOffset: number
  w: number; h: number
  isConstellationNode = false; constellationIndex = -1

  constructor(w: number, h: number) {
    this.w = w; this.h = h
    this.x = Math.random() * w; this.y = Math.random() * h
    this.vx = (Math.random() - 0.5) * PARTICLE_SPEED
    this.vy = (Math.random() - 0.5) * PARTICLE_SPEED
    this.radius = Math.random() * 1.7 + 0.3
    const colors = ['200,220,255','225,210,255','255,240,220','180,200,235','210,190,225','190,210,245','240,225,210','200,210,240']
    this.color = colors[Math.floor(Math.random() * colors.length)]
    this.opacity = Math.random() * 0.75 + 0.25
    this.twinkleSpeed = Math.random() * 0.025 + 0.004
    this.twinkleOffset = Math.random() * Math.PI * 2
  }

  update() {
    this.x += this.vx; this.y += this.vy
    const dx = this.x - mouse.x, dy = this.y - mouse.y
    const dist = Math.sqrt(dx * dx + dy * dy)
    if (dist < 140 && mouse.x > 0) {
      const angle = Math.atan2(dy, dx), force = (140 - dist) / 140
      this.vx += Math.cos(angle) * force * 0.17
      this.vy += Math.sin(angle) * force * 0.17
    }
    const speed = Math.sqrt(this.vx * this.vx + this.vy * this.vy)
    if (speed > 0.8) { this.vx = (this.vx / speed) * 0.8; this.vy = (this.vy / speed) * 0.8 }
    if (this.x < -40) this.x = this.w + 40
    if (this.x > this.w + 40) this.x = -40
    if (this.y < -40) this.y = this.h + 40
    if (this.y > this.h + 40) this.y = -40
  }

  draw(ctx: CanvasRenderingContext2D, time: number) {
    const twinkle = Math.sin(time * this.twinkleSpeed + this.twinkleOffset) * 0.3 + 0.7
    const op = this.opacity * twinkle
    ctx.beginPath(); ctx.arc(this.x, this.y, this.radius, 0, Math.PI * 2)
    ctx.fillStyle = `rgba(${this.color}, ${op})`; ctx.fill()
    if (this.radius > 1.2) {
      ctx.beginPath(); ctx.arc(this.x, this.y, this.radius * 2.5, 0, Math.PI * 2)
      ctx.fillStyle = `rgba(${this.color}, ${op * 0.1})`; ctx.fill()
    }
    if (this.isConstellationNode) {
      const glow = Math.sin(time * 0.003 + this.constellationIndex) * 0.4 + 0.6
      ctx.beginPath(); ctx.arc(this.x, this.y, this.radius * 3.8, 0, Math.PI * 2)
      ctx.fillStyle = `rgba(255,255,255, ${glow * 0.22})`; ctx.fill()
      ctx.strokeStyle = `rgba(255,255,255, ${glow * 0.45})`; ctx.lineWidth = 0.5
      const sl = this.radius * 5
      ctx.beginPath(); ctx.moveTo(this.x - sl, this.y); ctx.lineTo(this.x + sl, this.y)
      ctx.moveTo(this.x, this.y - sl); ctx.lineTo(this.x, this.y + sl); ctx.stroke()
    }
  }
}

function initConstellationNodes() {
  constellationNodes = []; constellationConnections = []
  const indices = new Set<number>()
  while (indices.size < CONSTELLATION_NODE_COUNT) indices.add(Math.floor(Math.random() * PARTICLE_COUNT))
  constellationNodes = Array.from(indices)
  constellationNodes.forEach((idx, i) => { particles[idx].isConstellationNode = true; particles[idx].constellationIndex = i })
  for (let i = 0; i < constellationNodes.length; i++) {
    for (let j = i + 1; j < constellationNodes.length; j++) {
      const pi = particles[constellationNodes[i]], pj = particles[constellationNodes[j]]
      const dx = pi.x - pj.x, dy = pi.y - pj.y
      const dist = Math.sqrt(dx * dx + dy * dy)
      if (dist < CONSTELLATION_CONNECT_DISTANCE && dist > 35) {
        constellationConnections.push({ i: constellationNodes[i], j: constellationNodes[j], phase: Math.random() * Math.PI * 2, speed: 0.0015 + Math.random() * 0.0035 })
      }
    }
  }
}

function drawConnections(ctx: CanvasRenderingContext2D, time: number) {
  // Regular connections — use squared distance to avoid Math.sqrt
  const maxDist2 = CONNECTION_DISTANCE * CONNECTION_DISTANCE
  ctx.lineWidth = 0.35
  ctx.strokeStyle = 'rgba(180,200,255,0.04)'
  ctx.beginPath()
  for (let i = 0; i < particles.length; i++) {
    const pi = particles[i]
    for (let j = i + 1; j < particles.length; j++) {
      const dx = pi.x - particles[j].x, dy = pi.y - particles[j].y
      if (dx * dx + dy * dy < maxDist2) {
        ctx.moveTo(pi.x, pi.y); ctx.lineTo(particles[j].x, particles[j].y)
      }
    }
  }
  ctx.stroke()

  // Constellation connections
  for (const conn of constellationConnections) {
    const pi = particles[conn.i], pj = particles[conn.j]
    const dx = pi.x - pj.x, dy = pi.y - pj.y
    const dist = Math.sqrt(dx * dx + dy * dy)
    if (dist < CONSTELLATION_CONNECT_DISTANCE * 1.3) {
      const pulse = Math.sin(time * conn.speed + conn.phase) * 0.5 + 0.5
      const opacity = pulse * 0.5 + 0.07
      const lw = pulse * 1.8 + 0.4
      ctx.beginPath(); ctx.moveTo(pi.x, pi.y); ctx.lineTo(pj.x, pj.y)
      ctx.strokeStyle = `rgba(220,230,255,${opacity})`; ctx.lineWidth = lw
      ctx.shadowColor = `rgba(180,200,255,${opacity * 0.7})`; ctx.shadowBlur = lw * 6
      ctx.stroke(); ctx.shadowColor = 'transparent'; ctx.shadowBlur = 0
    }
  }
}

function createExplosion(x: number, y: number, type: string) {
  const colorMap: Record<string, string[]> = {
    jupiter: ['#f5e6d3','#d4a373','#e8875b','#f0c080','#ffd4a0','#ffffff','#ff9944','#ff6600','#ffaa44'],
    saturn:  ['#f5ecd7','#e8d5a3','#d4b870','#c4a35a','#f0e0c0','#ffffff','#e0c890','#f8e8c0','#ccaa55'],
    earth:   ['#4a9eff','#2080e0','#ffffff','#5ab8ff','#0a60c0','#a0d8ff','#064090','#88ccff','#3a88ff'],
    mars:    ['#e87050','#d45a3a','#f08060','#c04028','#ff9060','#ffffff','#ff7040','#ff5500','#ffbb88'],
    moon:    ['#e8e8e8','#c8c8c8','#ffffff','#9a9a9a','#f0f0f0','#d0d0d0','#e0e0e0','#f8f8f8','#b0b0b0'],
  }
  const colors = colorMap[type] || ['#ffffff','#c8d8ff','#ffd4a0']

  // Main particles
  const count = 80 + Math.floor(Math.random() * 40)
  const eps: { x: number; y: number; vx: number; vy: number; life: number; decay: number; color: string; size: number }[] = []
  for (let i = 0; i < count; i++) {
    const angle = Math.random() * Math.PI * 2, speed = 1.5 + Math.random() * 13
    eps.push({ x, y, vx: Math.cos(angle) * speed, vy: Math.sin(angle) * speed, life: 1, decay: 0.005 + Math.random() * 0.022, color: colors[Math.floor(Math.random() * colors.length)], size: 1.5 + Math.random() * 5.5 })
  }

  // Spark particles — faster, smaller, brighter
  for (let i = 0; i < 25; i++) {
    const angle = Math.random() * Math.PI * 2, speed = 6 + Math.random() * 16
    eps.push({ x, y, vx: Math.cos(angle) * speed, vy: Math.sin(angle) * speed, life: 1, decay: 0.025 + Math.random() * 0.04, color: '#ffffff', size: 0.8 + Math.random() * 1.5 })
  }

  explosions.push({ x, y, particles: eps, shockwaveRadius: 8, shockwaveMaxRadius: 160 + Math.random() * 80, shockwaveLife: 1, shockwaveDecay: 0.022 })
  if (explosions.length > 6) explosions.shift()
}

function updateExplosions() {
  for (const exp of explosions) {
    exp.shockwaveLife -= exp.shockwaveDecay
    exp.shockwaveRadius += (exp.shockwaveMaxRadius - exp.shockwaveRadius) * 0.12 + 1.5
    for (const p of exp.particles) {
      p.x += p.vx; p.y += p.vy; p.vx *= 0.984; p.vy *= 0.984; p.vy += 0.04
      p.life -= p.decay; p.size *= 0.994
    }
    exp.particles = exp.particles.filter(p => p.life > 0)
  }
  explosions = explosions.filter(exp => exp.particles.length > 0 && exp.shockwaveLife > 0)
}

function drawExplosions(ctx: CanvasRenderingContext2D) {
  for (const exp of explosions) {
    // Outer shockwave ring
    if (exp.shockwaveLife > 0) {
      const sw1 = exp.shockwaveRadius
      ctx.beginPath(); ctx.arc(exp.x, exp.y, sw1, 0, Math.PI * 2)
      ctx.strokeStyle = `rgba(255,255,255,${exp.shockwaveLife * 0.7})`; ctx.lineWidth = 5 * exp.shockwaveLife
      ctx.shadowColor = `rgba(255,255,255,${exp.shockwaveLife})`; ctx.shadowBlur = 28 * exp.shockwaveLife; ctx.stroke()
      ctx.shadowColor = 'transparent'; ctx.shadowBlur = 0

      // Inner shockwave
      const sw2 = sw1 * 0.65
      ctx.beginPath(); ctx.arc(exp.x, exp.y, sw2, 0, Math.PI * 2)
      ctx.strokeStyle = `rgba(255,220,180,${exp.shockwaveLife * 0.5})`; ctx.lineWidth = 2.5 * exp.shockwaveLife; ctx.stroke()

      // Core flash
      const sw3 = sw1 * 0.35
      ctx.beginPath(); ctx.arc(exp.x, exp.y, sw3, 0, Math.PI * 2)
      ctx.fillStyle = `rgba(255,255,255,${exp.shockwaveLife * 0.25})`; ctx.fill()
    }

    for (const p of exp.particles) {
      const alphaHex = Math.floor(p.life * 255).toString(16).padStart(2, '0')
      // Main particle
      ctx.beginPath(); ctx.arc(p.x, p.y, p.size, 0, Math.PI * 2)
      ctx.fillStyle = p.color + alphaHex; ctx.fill()
      // Outer glow
      ctx.beginPath(); ctx.arc(p.x, p.y, p.size * 2.8, 0, Math.PI * 2)
      ctx.fillStyle = p.color + '44'; ctx.fill()
      // Inner bright spot
      if (p.size > 2) {
        ctx.beginPath(); ctx.arc(p.x, p.y, p.size * 0.4, 0, Math.PI * 2)
        ctx.fillStyle = '#ffffff' + alphaHex; ctx.fill()
      }
    }
  }
}

function getPlanetCenter(el: Element) {
  const rect = el.getBoundingClientRect()
  return { x: rect.left + rect.width / 2, y: rect.top + rect.height / 2, radius: Math.max(rect.width, rect.height) / 2 }
}

function onDocClick(e: MouseEvent) {
  if (!containerRef.value) return
  for (const sel of planetSelectors) {
    const el = containerRef.value.querySelector('.' + sel)
    if (!el) continue
    const pos = getPlanetCenter(el)
    const dx = e.clientX - pos.x, dy = e.clientY - pos.y
    if (Math.sqrt(dx * dx + dy * dy) < pos.radius * 1.25) {
      createExplosion(pos.x, pos.y, sel)
      mouse.clickX = pos.x; mouse.clickY = pos.y; mouse.clicked = true
      setTimeout(() => { mouse.clicked = false }, 300)
      break
    }
  }
}

function onTouchStart(e: TouchEvent) {
  if (!containerRef.value || e.touches.length === 0) return
  const cx = e.touches[0].clientX, cy = e.touches[0].clientY
  for (const sel of planetSelectors) {
    const el = containerRef.value.querySelector('.' + sel)
    if (!el) continue
    const pos = getPlanetCenter(el)
    const dx = cx - pos.x, dy = cy - pos.y
    if (Math.sqrt(dx * dx + dy * dy) < pos.radius * 1.3) {
      createExplosion(pos.x, pos.y, sel)
      mouse.clickX = pos.x; mouse.clickY = pos.y; mouse.clicked = true
      setTimeout(() => { mouse.clicked = false }, 300)
      break
    }
  }
}

function onMouseMove(e: MouseEvent) {
  mouse.x = e.clientX; mouse.y = e.clientY
  if (glowRef.value) { glowRef.value.style.left = e.clientX + 'px'; glowRef.value.style.top = e.clientY + 'px'; glowRef.value.style.opacity = '1' }
  const moveX = (e.clientX - window.innerWidth / 2) / window.innerWidth
  const moveY = (e.clientY - window.innerHeight / 2) / window.innerHeight
  if (containerRef.value) {
    containerRef.value.querySelectorAll('.nebula').forEach((el, i) => { (el as HTMLElement).style.transform = `translate(${moveX * (12 + i * 3)}px, ${moveY * (12 + i * 3)}px)` })
    containerRef.value.querySelectorAll('.planet').forEach((el, i) => { (el as HTMLElement).style.transform = `translate(${moveX * (20 + i * 6)}px, ${moveY * (20 + i * 6)}px)` })
    containerRef.value.querySelectorAll('.star-dot').forEach((el, i) => { (el as HTMLElement).style.transform = `translate(${moveX * (10 + i * 2)}px, ${moveY * (10 + i * 2)}px)` })
  }
}

function onMouseLeave() {
  mouse.x = -1000; mouse.y = -1000
  if (glowRef.value) glowRef.value.style.opacity = '0'
  if (containerRef.value) {
    containerRef.value.querySelectorAll('.nebula, .planet, .star-dot').forEach(el => { (el as HTMLElement).style.transform = 'translate(0px, 0px)' })
  }
}

function onResize() {
  const canvas = canvasRef.value
  if (!canvas) return
  const w = window.innerWidth, h = window.innerHeight
  canvas.width = w; canvas.height = h
  if (particles.length > 0) {
    particles.forEach(p => { p.w = w; p.h = h })
    particles.forEach(p => { p.isConstellationNode = false; p.constellationIndex = -1 })
    initConstellationNodes()
  }
}

onMounted(() => {
  alive = true
  const canvas = canvasRef.value!
  const ctx = canvas.getContext('2d')!

  onResize()
  window.addEventListener('resize', onResize)
  document.addEventListener('mousemove', onMouseMove)
  document.addEventListener('mouseleave', onMouseLeave)
  document.addEventListener('click', onDocClick)
  document.addEventListener('touchstart', onTouchStart, { passive: true })

  particles = Array.from({ length: PARTICLE_COUNT }, () => new Particle(canvas.width, canvas.height))
  initConstellationNodes()

  function animate(time: number) {
    if (!alive) return
    ctx.clearRect(0, 0, canvas.width, canvas.height)
    for (const p of particles) { p.update(); p.draw(ctx, time) }
    drawConnections(ctx, time)
    updateExplosions(); drawExplosions(ctx)
    if (mouse.clicked && mouse.clickX >= 0) {
      const grad = ctx.createRadialGradient(mouse.clickX, mouse.clickY, 0, mouse.clickX, mouse.clickY, 90)
      grad.addColorStop(0, 'rgba(255,255,255,0.85)'); grad.addColorStop(0.3, 'rgba(255,255,255,0.35)'); grad.addColorStop(1, 'rgba(255,255,255,0)')
      ctx.fillStyle = grad; ctx.beginPath(); ctx.arc(mouse.clickX, mouse.clickY, 90, 0, Math.PI * 2); ctx.fill()
    }
    animId = requestAnimationFrame(animate)
  }
  animId = requestAnimationFrame(animate)
})

onUnmounted(() => {
  alive = false
  cancelAnimationFrame(animId)
  window.removeEventListener('resize', onResize)
  document.removeEventListener('mousemove', onMouseMove)
  document.removeEventListener('mouseleave', onMouseLeave)
  document.removeEventListener('click', onDocClick)
  document.removeEventListener('touchstart', onTouchStart)
})
</script>

<style scoped lang="scss">
.space-bg {
  position: absolute;
  inset: 0;
  overflow: hidden;
  pointer-events: none;
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
  opacity: 0.28;
  will-change: transform;
  animation: floatNebula 20s infinite alternate ease-in-out;
  mix-blend-mode: screen;
}
.nebula-1 {
  width: 70vw; height: 70vw; max-width: 850px; max-height: 850px;
  background: radial-gradient(circle, #1a1a35 0%, transparent 70%);
  top: -12%; left: -8%; animation-duration: 22s;
}
.nebula-2 {
  width: 75vw; height: 75vw; max-width: 950px; max-height: 950px;
  background: radial-gradient(circle, #25183a 0%, transparent 70%);
  bottom: -18%; right: -10%; animation-duration: 24s; animation-delay: -6s;
}
.nebula-3 {
  width: 55vw; height: 55vw; max-width: 650px; max-height: 650px;
  background: radial-gradient(circle, #0e1a2a 0%, transparent 70%);
  top: 48%; left: 38%; opacity: 0.22; animation-duration: 28s; animation-delay: -10s;
}
@keyframes floatNebula {
  0% { transform: translate(0, 0) scale(1); }
  33% { transform: translate(20px, -35px) scale(1.05); }
  66% { transform: translate(-15px, 22px) scale(0.95); }
  100% { transform: translate(12px, -18px) scale(1.02); }
}

// Planets
.planet {
  position: absolute;
  border-radius: 50%;
  will-change: transform;
  animation: floatPlanet 24s infinite alternate ease-in-out;
  pointer-events: auto;
  cursor: pointer;
  transition: filter 0.4s ease, transform 0.4s ease;

  &:hover {
    filter: brightness(1.3);
    transform: scale(1.12) !important;
    z-index: 10;
  }
}

// Jupiter — massive gas giant with cream/tan bands and Great Red Spot
.jupiter {
  width: 150px; height: 150px; top: 8%; right: 10%;
  background:
    radial-gradient(circle at 32% 28%, rgba(255,255,240,0.4) 0%, rgba(255,250,230,0.1) 25%, transparent 45%),
    radial-gradient(ellipse 20px 13px at 56% 64%, rgba(180,70,35,0.55) 0%, rgba(160,55,25,0.25) 65%, transparent 100%),
    repeating-linear-gradient(
      0deg,
      #e8d8b8 0px, #dcc8a0 2px,
      #c8b48a 4px, #b09868 6px,
      #d4c098 8px, #e0cca0 10px,
      #c8a878 12px, #b09060 14px,
      #d8c498 16px, #e8d8b8 18px,
      #c4a870 20px, #b89868 22px,
      #dcc8a0 24px, #c8b088 26px,
      #a88858 28px, #d0bc90 30px,
      #e0cca0 32px, #c8a878 34px,
      #b89868 36px, #d8c498 38px
    );
  box-shadow:
    0 0 50px rgba(200, 170, 110, 0.45),
    0 0 100px rgba(200, 170, 110, 0.15),
    0 0 150px rgba(200, 170, 110, 0.06),
    inset -14px -10px 32px rgba(60,40,20,0.7),
    inset 8px 6px 22px rgba(255,240,210,0.18);
  animation-duration: 28s; animation-delay: -3s;

  &::before {
    content: '';
    position: absolute; inset: -4px;
    border-radius: 50%;
    background: radial-gradient(circle at 50% 50%, transparent 44%, rgba(210,180,120,0.1) 58%, rgba(210,180,120,0.03) 72%, transparent 100%);
    pointer-events: none;
  }
}

// Saturn — golden gas giant with prominent multi-band ring
.saturn {
  width: 115px; height: 115px; top: 14%; left: 6%;
  background:
    radial-gradient(circle at 30% 25%, rgba(255,255,255,0.3) 0%, rgba(255,255,255,0.06) 30%, transparent 50%),
    repeating-linear-gradient(
      0deg,
      #e8d5a3 0px, #d4b870 3px, #c4a855 5px,
      #b8943a 7px, #d4b870 10px, #e8d8a0 12px,
      #f0e0c0 14px, #d4b870 16px, #c4a855 18px,
      #e0c880 20px, #d4b870 23px
    );
  box-shadow:
    0 0 45px rgba(220, 190, 140, 0.4),
    0 0 90px rgba(220, 190, 140, 0.15),
    inset -10px -8px 25px rgba(0,0,0,0.65),
    inset 5px 5px 15px rgba(255,230,180,0.15);
  animation-duration: 26s; animation-delay: -5s;

  &::before {
    content: '';
    position: absolute; inset: -3px;
    border-radius: 50%;
    background: radial-gradient(circle at 50% 50%, transparent 42%, rgba(220,190,140,0.1) 58%, rgba(220,190,140,0.03) 72%, transparent 100%);
    pointer-events: none;
  }
}

.saturn-ring {
  position: absolute;
  top: 50%; left: 50%;
  width: 195px; height: 55px;
  transform: translate(-50%, -50%) rotateX(75deg) rotateZ(-15deg);
  border-radius: 50%;
  background: transparent;
  box-shadow:
    0 0 0 2px rgba(230,200,150,0.08),
    0 0 0 5px rgba(210,185,140,0.3),
    0 0 0 7px rgba(190,160,110,0.12),
    0 0 0 10px rgba(230,200,150,0.35),
    0 0 0 12px rgba(210,185,140,0.15),
    0 0 0 15px rgba(230,200,150,0.25),
    0 0 0 17px rgba(200,170,120,0.08),
    0 0 0 20px rgba(230,200,150,0.12);
  pointer-events: none;
}

// Earth — deep blue ocean with white cloud wisps and atmosphere halo
.earth {
  width: 90px; height: 90px; bottom: 18%; left: 13%;
  background:
    radial-gradient(circle at 30% 26%, rgba(255,255,255,0.5) 0%, rgba(200,230,255,0.15) 20%, transparent 40%),
    radial-gradient(ellipse 40px 3px at 35% 38%, rgba(255,255,255,0.35) 0%, transparent 100%),
    radial-gradient(ellipse 30px 2px at 55% 55%, rgba(255,255,255,0.25) 0%, transparent 100%),
    radial-gradient(ellipse 22px 2px at 45% 70%, rgba(255,255,255,0.2) 0%, transparent 100%),
    radial-gradient(circle at 40% 8%, rgba(230,245,255,0.6) 0%, rgba(200,230,255,0.2) 10%, transparent 20%),
    radial-gradient(circle at 60% 92%, rgba(230,245,255,0.5) 0%, rgba(200,230,255,0.15) 10%, transparent 18%),
    radial-gradient(circle at 38% 38%, #4a9eff 0%, #2080e0 25%, #0a60c0 50%, #064090 72%, #022860 100%);
  box-shadow:
    0 0 40px rgba(50, 140, 255, 0.55),
    0 0 80px rgba(50, 140, 255, 0.2),
    0 0 120px rgba(50, 140, 255, 0.08),
    inset -9px -7px 22px rgba(0,10,40,0.6),
    inset 5px 5px 14px rgba(160,210,255,0.2);
  animation-duration: 22s; animation-delay: -8s;

  &::before {
    content: '';
    position: absolute; inset: -7px;
    border-radius: 50%;
    background: radial-gradient(circle at 50% 50%, transparent 40%, rgba(60,150,255,0.18) 54%, rgba(60,150,255,0.05) 68%, transparent 100%);
    pointer-events: none;
  }
}

// Mars — rust red with polar cap, dark terrain features
.mars {
  width: 58px; height: 58px; top: 46%; right: 8%;
  background:
    radial-gradient(circle at 28% 25%, rgba(255,220,180,0.35) 0%, transparent 40%),
    radial-gradient(ellipse 10px 7px at 48% 38%, rgba(80,30,10,0.5) 0%, rgba(60,20,5,0.3) 70%, transparent 100%),
    radial-gradient(ellipse 8px 6px at 22% 55%, rgba(60,20,5,0.4) 0%, transparent 100%),
    radial-gradient(ellipse 6px 5px at 68% 60%, rgba(90,35,10,0.35) 0%, transparent 100%),
    radial-gradient(ellipse 14px 5px at 50% 8%, rgba(240,230,220,0.5) 0%, rgba(220,210,200,0.2) 60%, transparent 100%),
    radial-gradient(circle, #e07050 0%, #c04530 35%, #8a2510 65%, #4a1005 100%);
  box-shadow:
    0 0 30px rgba(220, 100, 60, 0.45),
    0 0 60px rgba(220, 100, 60, 0.15),
    inset -6px -5px 14px rgba(0,0,0,0.6),
    inset 4px 3px 10px rgba(255,180,140,0.2);
  animation-duration: 20s; animation-delay: -9s;
}

// Moon — gray with detailed craters, bright highlight
.moon {
  width: 62px; height: 62px; top: 54%; left: 26%;
  background:
    radial-gradient(circle at 28% 25%, rgba(255,255,255,0.35) 0%, transparent 40%),
    radial-gradient(circle 5px at 50% 42%, rgba(40,40,40,0.55) 0%, rgba(40,40,40,0.55) 60%, transparent 100%),
    radial-gradient(circle 4px at 28% 62%, rgba(50,50,50,0.45) 0%, rgba(50,50,50,0.45) 60%, transparent 100%),
    radial-gradient(circle 3px at 68% 32%, rgba(35,35,35,0.4) 0%, rgba(35,35,35,0.4) 60%, transparent 100%),
    radial-gradient(circle 3px at 42% 78%, rgba(55,55,55,0.35) 0%, rgba(55,55,55,0.35) 60%, transparent 100%),
    radial-gradient(circle 6px at 72% 58%, rgba(45,45,45,0.3) 0%, rgba(45,45,45,0.3) 60%, transparent 100%),
    radial-gradient(circle at 45% 45%, #dcdcdc 0%, #b8b8b8 30%, #909090 60%, #585858 85%, #383838 100%);
  box-shadow:
    0 0 30px rgba(200,200,200,0.35),
    0 0 60px rgba(200,200,200,0.1),
    inset -5px -4px 12px rgba(0,0,0,0.65),
    inset 3px 3px 8px rgba(255,255,255,0.15);
  animation-duration: 20s; animation-delay: -12s;
}

// Orbital rings — planet-specific styling
.jupiter::after, .saturn::after, .earth::after, .moon::after, .mars::after {
  content: '';
  position: absolute; top: 50%; left: 50%;
  transform: translate(-50%, -50%) rotateX(75deg) rotateY(20deg);
  border-radius: 50%;
  border: 2px solid rgba(255,255,255,0.15);
  box-shadow: 0 0 12px rgba(255,255,255,0.1);
  animation: ringPulse 8s infinite alternate ease-in-out;
  pointer-events: none;
}
.jupiter::after { width: 190px; height: 190px; border-width: 2.5px; border-color: rgba(210,180,140,0.3); box-shadow: 0 0 25px rgba(210,180,140,0.2), 0 0 50px rgba(210,180,140,0.06); }
.earth::after { width: 125px; height: 125px; border-width: 1.5px; border-color: rgba(100,180,255,0.35); box-shadow: 0 0 22px rgba(0,180,255,0.25), 0 0 45px rgba(0,180,255,0.08); }
.mars::after { width: 85px; height: 85px; border-width: 1px; border-style: dashed; border-color: rgba(220,150,120,0.3); box-shadow: 0 0 12px rgba(200,120,80,0.18); animation-duration: 7s; }
.moon::after { width: 88px; height: 88px; border-width: 1px; border-style: dashed; border-color: rgba(200,200,220,0.3); box-shadow: 0 0 12px rgba(180,180,210,0.18); animation-duration: 5s; }

@keyframes floatPlanet {
  0% { transform: translate(0, 0) scale(1); }
  33% { transform: translate(25px, -45px) scale(1.03); }
  66% { transform: translate(-20px, 30px) scale(0.97); }
  100% { transform: translate(15px, -25px) scale(1.01); }
}
@keyframes ringPulse {
  0% { opacity: 0.5; transform: translate(-50%, -50%) rotateX(75deg) rotateY(20deg) scale(1); }
  100% { opacity: 0.85; transform: translate(-50%, -50%) rotateX(75deg) rotateY(20deg) scale(1.02); }
}

// Star dots
.star-dot {
  position: absolute; border-radius: 50%; background: #b0e0ff;
  box-shadow: 0 0 12px #7bb3d9, 0 0 25px rgba(100,150,255,0.5);
  animation: twinkle 10s infinite alternate ease-in-out;
}
.dot-1 { width: 8px; height: 8px; top: 28%; left: 18%; background: #c4d2ff; }
.dot-2 { width: 5px; height: 5px; top: 72%; right: 22%; background: #ffe0b0; animation-duration: 12s; animation-delay: -3s; }
.dot-3 { width: 6px; height: 6px; bottom: 28%; left: 52%; background: #c9e4ff; animation-duration: 14s; animation-delay: -5s; }
@keyframes twinkle {
  0% { opacity: 0.55; transform: scale(1); }
  100% { opacity: 1; transform: scale(1.6); }
}

// Star trails
.star-trail {
  position: absolute; bottom: 0; left: 0; width: 100%; height: 2px;
  background: linear-gradient(90deg, transparent, #3a5a8a, #8098b8, transparent);
  animation: slideLine 8s linear infinite; filter: blur(2px); opacity: 0.35;
}
.star-trail-2 { top: 0; bottom: auto; animation-duration: 12s; background: linear-gradient(90deg, transparent, #4a3568, #7a6a9a, transparent); animation-delay: -4s; }
@keyframes slideLine {
  0% { transform: translateX(-105%); }
  100% { transform: translateX(105%); }
}

// Cursor glow
.cursor-glow {
  position: absolute; width: 420px; height: 420px; border-radius: 50%;
  background: radial-gradient(circle, rgba(140,160,220,0.1) 0%, transparent 70%);
  filter: blur(65px); pointer-events: none; z-index: 2;
  transform: translate(-50%, -50%); transition: opacity 0.3s ease; opacity: 0;
}

@media (max-width: 768px) {
  .jupiter { width: 95px; height: 95px; }
  .saturn { width: 75px; height: 75px; }
  .earth { width: 60px; height: 60px; }
  .mars { width: 40px; height: 40px; }
  .moon { width: 42px; height: 42px; }
  .saturn-ring { width: 130px; height: 38px; }
  .cursor-glow { width: 180px; height: 180px; }
  .jupiter::after { width: 125px; height: 125px; }
  .earth::after { width: 85px; height: 85px; }
  .mars::after { width: 58px; height: 58px; }
  .moon::after { width: 60px; height: 60px; }
}
</style>
