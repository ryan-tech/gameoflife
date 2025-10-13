import type { FC } from 'react';
import Cell from './Cell';
import './Grid.css';

interface GridProps {
    rows: number;
    cols: number;
}

const Grid: FC<GridProps> = ({ rows = 20, cols = 20 }) => {
    return (
        <div
            className="grid"
            style={{
                gridTemplateColumns: `repeat(${cols}, 20px)`,
            }}
        >
            {Array.from({ length: rows * cols }).map((_, i) => (
                <Cell key={i} onClick={() => {}} />
            ))}
        </div>
    );
};

export default Grid;
