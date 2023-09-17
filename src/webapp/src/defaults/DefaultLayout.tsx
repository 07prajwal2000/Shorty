import React from 'react'
import Navbar from '../components/Navbar';

interface LayoutProps {
  children: React.ReactNode | React.ReactNode[];
}

const DefaultLayout: React.FC<LayoutProps> = ({children}) => {
  return (
    <>
      <Navbar />
      <main>{children}</main>
    </>
  )
}

export default DefaultLayout