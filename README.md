# Quadlab — Podman Quadlet Homelab Infrastructure

Quadlab is a personal homelab project built to explore **declarative, reproducible infrastructure** using Podman, Quadlets, and systemd — without relying on Docker daemons or orchestration frameworks.

The goal of this project is to model *production-style service management* in a small, understandable environment while keeping configuration explicit, auditable, and easy to reproduce.

Quadlab is not intended to be a replacement for container orchestration tools like Kubernetes, but to push the limits of reproducibility and declarative infrastructure with the tools that exist on Fedora Atomic spins

## Why Quadlab?

Most homelabs grow organically:
- ad-hoc Docker Compose files
- manual systemd services
- configuration scattered across the host

Quadlab experiments with a different approach:
- **rootless containers**
- **systemd-managed lifecycles**
- **declarative service definitions**
- minimal host configuration

By using **Podman Quadlets**, container services become first-class systemd units, inheriting all the benefits of systemd:
- dependency ordering
- restart policies
- logging via journald
- timers instead of cron

## Why Fedora Atomic?

Quadlab is intentionally built on a Fedora Atomic host, where the base operating system is **immutable** and managed as an image rather than a mutable package set.

The goal is to push reproducibility as far as practical:

- Minimize host-level configuration
- Avoid configuration drift
- Treat the OS as replaceable, not precious
- Keep all service logic in declarative definitions

In this model:
- The host provides a stable, minimal substrate
- Containers hold application logic
- systemd manages lifecycle and scheduling
- Rebuilding or replacing the host should require minimal effort

## Core ideas

- **No daemon dependency**  
  Podman runs containers directly — systemd handles orchestration.

- **Declarative infrastructure**  
  Services are defined in `.container` and `.network` files, not imperative scripts.

- **Rootless by default**  
  Containers run as an unprivileged user, reducing blast radius.

- **Reproducibility**  
  A new machine can reproduce the setup by syncing configuration and reloading systemd.

## What’s included

This repo represents a working lab environment that includes:

- Reverse proxy (Caddy)
- Dynamic DNS updater (custom Go service)
- Media services (Jellyfin)
- Application services (e.g. FoundryVTT)
- Local DNS (CoreDNS)
- Custom systemd timers for scheduled jobs
- Container networking via Podman networks

Services are managed using Quadlet `.container` files and user-level systemd units.

## Repository layout

- `containers/` — Quadlet `.container` definitions
- `networks/` — Podman network definitions
- `timers/` — systemd timers for scheduled tasks
- `scripts/` — helper scripts and tooling
- `docs/` (optional) — architecture notes and diagrams

## How it works

1. Quadlet files are placed in:

~/.config/containers/systemd/

2. systemd loads them as user services
3. Podman creates and runs containers automatically
4. Logs and lifecycle events are managed by systemd/journald
5. systemd timers replace cron for scheduled container jobs


## Why Quadlets instead of Docker Compose?

Docker Compose works well, but:
- relies on a long-running daemon
- introduces an extra control layer
- separates container lifecycle from the OS init system

Quadlets let systemd do what it already does well:
- start services
- restart them
- order dependencies
- schedule jobs

This keeps the system simpler and easier to reason about.
