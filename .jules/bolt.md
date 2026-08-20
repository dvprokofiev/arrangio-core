## 2024-08-14 - Inner Loop Branch Prediction
**Learning:** Branch checking inside the innermost loop of a 3D grid traversal is a significant bottleneck. When checking if a cell coordinate is out-of-bounds (`x < 0 || x >= size...`), it hurts branch prediction and forces unnecessary iterations for entities partially out of bounds.
**Action:** Hoist bounds checking outside the loops by pre-clamping the min/max bounds using `max(0, ...)` and `min(size-1, ...)`.
## 2024-08-14 - Loop Ordering and Index Calculation in 3D Grid
**Learning:** The spatial grid insertion and query performance is sensitive to the nested loop order and redundant calculations within the inner loops. The previous `x, y, z` nested loop required calling `getIndex` which multiplied inner variables repeatedly. Additionally, traversing slices sequentially is usually better for cache performance.
**Action:** Order 3D grid traversal loops as `z, y, x` (outer to inner) to allow precomputing `zOffset` and `yOffset`, reducing the inner loop's address calculation to a simple addition `x + yOffset`.
## 2024-08-16 - Hoisting Virtual Method Calls in Spatial Loops
**Learning:** Calling virtual interface methods like `b.Shape.Bounds()` or even wrapping method calls like `b.ContainsPoint(wx, wy, wz)` inside the innermost loop of a spatial intersection test (like `a.Shape.ForEachPoint`) introduces significant overhead and redundant calculations. Pre-calculating world bounds and coordinate differentials outside the loop avoids these redundant operations.
**Action:** When implementing point-wise iteration operations, hoist bounds calculations, anchor transformations, and method delegations out of the inner loop and instead use precomputed world-space coordinates and inlined checks.
## 2025-02-12 - Integer Overflow in World Bounds Calculation
**Learning:** Subtracting world coordinates (`int64`) and casting immediately to `int16` can cause integer overflow if objects are vastly distant (e.g., >32767 apart). This causes phantom collisions (false positives).
**Action:** When precomputing relative local bounds or offsets, perform the subtraction in `int64` world space first, then cast the result to `int16`, or just use the resulting `int64` for bounds checking to avoid overflow vulnerabilities.

## 2025-02-12 - Devirtualization vs Inlining in Go
**Learning:** Creating a function pointer to an interface method (`checkFn = bShape.Contains`) still acts as an indirect call, incurs a closure allocation (if the receiver is bound), and prevents compiler inlining. It is generally not faster than a direct interface call unless the interface dispatch itself is a major bottleneck (which is rarely true compared to the inlining benefits).
**Action:** To truly devirtualize in Go, type-switch *outside* the hot loop, and duplicate the inner loop for each concrete type (loop unswitching). This allows the compiler to fully inline the concrete method calls (like `geometry.Box.Contains`) and avoid closure allocations.
