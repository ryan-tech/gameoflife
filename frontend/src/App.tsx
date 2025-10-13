import { useState, useRef } from 'react'
import './App.css'
import Grid from './components/Grid'
import Controls from './components/Controls'

function App() {
  const [isPlaying, setIsPlaying] = useState(false)
  const [_, setGeneration] = useState(0)
  const intervalRef = useRef<NodeJS.Timeout | null>(null)

  const handlePlay = () => {
    setIsPlaying(true)
    // Start the game loop
    intervalRef.current = setInterval(() => {
      // Game logic here - advance one generation
      setGeneration(prev => prev + 1)
    }, 500) // 500ms between generations
  }

  const handlePause = () => {
    setIsPlaying(false)
    if (intervalRef.current) {
      clearInterval(intervalRef.current)
      intervalRef.current = null
    }
  }

  const handleStep = () => {
    // Advance one generation manually
    setGeneration(prev => prev + 1)
  }

  const handleReset = () => {
    setIsPlaying(false)
    setGeneration(0)
    if (intervalRef.current) {
      clearInterval(intervalRef.current)
      intervalRef.current = null
    }
    // Reset grid state here
  }

  return (
    <>
      <Grid rows={5} cols={5} />
      <Controls 
        onPlay={handlePlay}
        onPause={handlePause}
        onStep={handleStep}
        onReset={handleReset}
        isPlaying={isPlaying}
      />
    </>
  )
}

export default App
