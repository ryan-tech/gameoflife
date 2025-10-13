import ControlButton from './ControlButton';

interface ControlsProps {
    onPlay: () => void;
    onPause: () => void;
    onStep: () => void;
    onReset: () => void;
    isPlaying: boolean;
}

const Controls = ({ onPlay, onPause, onStep, onReset, isPlaying }: ControlsProps) => {
    return (
        <div id="controls-container">
            <ControlButton 
                text={isPlaying ? "Playing..." : "Play"} 
                onClick={onPlay}
                disabled={isPlaying}
            />
            <ControlButton 
                text="Pause" 
                onClick={onPause}
                disabled={!isPlaying}
            />
            <ControlButton 
                text="Step" 
                onClick={onStep}
                disabled={isPlaying}
            />
            <ControlButton 
                text="Reset" 
                onClick={onReset}
            />
        </div>
    )
}

export default Controls;