Game of Life Basics

As mentioned earlier, the Game of Life consists of an (infinite) grid where all cells at any given timestep are either alive or dead.

The rules are very simple. If a cell is alive at one timestep, it remains alive if two or three neighbors (determined as the four directly adjacent cells along with the four additional diagonally-adjacent cells) are alive; otherwise, it dies (as if by over-population). Conversely, if a cell is dead at some timestep, it becomes alive at the next timestep if exactly three of its neighbors are alive, otherwise it remains dead. These rules turn out to be exceptionally powerful. 

This is a project to showcase the integration of Go and Three.js.
