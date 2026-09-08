# Spring animation

## Integration

Create a spring with `harmonica.NewSpring(harmonica.FPS(fps), angularFrequency, dampingRatio)`. `Update(position, velocity, target)` returns both the next position and next velocity. Preserve both across messages; discarding velocity changes the simulation.

The timestep describes the simulation, not the renderer. Match the configured FPS to the command schedule, or intentionally manage variable elapsed time with a model appropriate to the application. Do not assume `WithFPS` makes every animation tick arrive at that rate under load.

## Tick ownership

A Bubble Tea tick command runs outside Update and returns a tick message. The handler performs the small numerical step, updates state, and returns another tick only if still active. Do not copy a standalone `for { time.Sleep(...) }` loop into Update or View.

When a new target arrives, update the existing animation state and ensure only one tick chain remains active. A generation counter lets the model discard old tick messages after cancellation/restart. When reduced motion is enabled, set position to target, velocity to zero, and stop rescheduling.

## Tuning and testing

Under-damped motion overshoots; critical damping reaches the target without oscillation; over-damping is slower. Choose motion that fits the UI and avoid infinite tiny movement. Use tolerances tied to terminal cell scale and desired responsiveness, not unexplained constants copied from an unrelated example.

Test the numerical transition directly with fixed steps. Verify convergence, bounded overshoot when relevant, retargeting, stale ticks, and stopping commands. Convert the final position to a terminal coordinate deliberately; floating point precision is not visible screen precision.

Consult the exact pinned APIs and examples in [sources](sources.md).
