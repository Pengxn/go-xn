import { useState } from 'react'

import Avatar from './components/Avatar.jsx'

const App = () => {
    const [count, setCount] = useState(0)

    const handleClick = () => setCount(count + 1)

    return (
        <>
            <div className='text-center w-full h-full flex flex-col justify-center items-center'>
                <h1 className='text-center text-3xl font-bold'>
                    Hello world!
                </h1>
                <div>
                    <div onClick={handleClick}>
                        Click avatar to add count number: {count}
                    </div>
                </div>
                <div onClick={handleClick}>
                    <Avatar email='i@fengyj.cn' />
                </div>
            </div>
        </>
    )
}

export default App
