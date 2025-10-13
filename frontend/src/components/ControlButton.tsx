// import './ControlButton.css';

interface ControlButtonProps {
    text: string;
    onClick: () => void;
    disabled?: boolean;
}

const ControlButton = ({ text, onClick, disabled = false }: ControlButtonProps) => {  
    return (
        <button onClick={onClick} disabled={disabled}>
            {text}
        </button>
    )
}

export default ControlButton;