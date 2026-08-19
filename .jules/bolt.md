## 2024-08-14 - Inner Loop Branch Prediction
**Learning:** Branch checking inside the innermost loop of a 3D grid traversal is a significant bottleneck. When checking if a cell coordinate is out-of-bounds (`x < 0 || x >= size...`), it hurts branch prediction and forces unnecessary iterations for entities partially out of bounds.
**Action:** Hoist bounds checking outside the loops by pre-clamping the min/max bounds using `max(0, ...)` and `min(size-1, ...)`.
## 2024-08-14 - Loop Ordering and Index Calculation in 3D Grid
**Learning:** The spatial grid insertion and query performance is sensitive to the nested loop order and redundant calculations within the inner loops. The previous `x, y, z` nested loop required calling `getIndex` which multiplied inner variables repeatedly. Additionally, traversing slices sequentially is usually better for cache performance.
**Action:** Order 3D grid traversal loops as `z, y, x` (outer to inner) to allow precomputing `zOffset` and `yOffset`, reducing the inner loop's address calculation to a simple addition `x + yOffset`.
## 2024-08-16 - Hoisting Virtual Method Calls in Spatial Loops
**Learning:** Calling virtual interface methods like `b.Shape.Bounds()` or even wrapping method calls like `b.ContainsPoint(wx, wy, wz)` inside the innermost loop of a spatial intersection test (like `a.Shape.ForEachPoint`) introduces significant overhead and redundant calculations. Pre-calculating world bounds and coordinate differentials outside the loop avoids these redundant operations.
**Action:** When implementing point-wise iteration operations, hoist bounds calculations, anchor transformations, and method delegations out of the inner loop and instead use precomputed world-space coordinates and inlined checks.
## 2024-08-16 - Type Asserts Over Interfaces in Tight Loops
**Learning:** Even when delegating the `Contains` check out of a wrapper function, an interface method call inside a tight loop (`ForEachPoint`) adds measurable overhead. Extracting the underlying type with a type switch and calling its concrete method can improve performance significantly.
**Action:** In high-frequency loops where interface methods are called, use type switches for common shape types (e.g., Box, VoxelShape) to perform direct, non-virtual method calls.
