# Conway's Game of Life

## Introduction
Conway's Game Of Life is a simulation of life in a 2 dimensional array. Each cell in the grid can be in an alive or dead state. 
The rules for cell evolution are: a live cell with fewer than two live neighbors dies (underpopulation), a live cell with two or three live neighbors survives, a live cell with more than three live neighbors dies (overpopulation), and a dead cell with exactly three live neighbors becomes alive (reproduction).

This project uses Golang to simulate the game logic and gRPC to expose API functionality.