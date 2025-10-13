import type { FC } from 'react';
import { useState } from 'react';
import './Cell.css';


interface CellProps {
    key: number;
    onClick: () => void;
}

const Cell: FC<CellProps> = ({ key, onClick }) => {
    const [aliveState, setAliveState] = useState(false);
    return <div
          key={key}
          onClick={() => {
              onClick();
              setAliveState(!aliveState);
          }}
          className={`cell ${aliveState ? 'alive' : 'dead'}`}
        
    />
}

export default Cell;