import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import RecoilProvider from './provider/recoilProvider.tsx'
import { Toaster } from 'sonner'
import { RouterProvider } from 'react-router-dom'
import { Routes } from './routes/routes.tsx'

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <RecoilProvider>
    <RouterProvider router={Routes} />
    <Toaster richColors/>
    </RecoilProvider>
  </StrictMode>,
)
