import Header from './Header'
import Footer from './Footer'

const MainLayout = ({ children }: { children: React.ReactNode }) => {
  return (
    <div>
      <Header />
      <div className="container">
        {children}
      </div>
      <Footer />
    </div>
  )
}

export default MainLayout