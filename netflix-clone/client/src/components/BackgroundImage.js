import React from 'react'
import styled from 'styled-components'

const BackgroundImage = () => {
  return (
    <BackgroundContainer>
      <img src = 'https://res.cloudinary.com/ehizeex-shop/image/upload/v1668265236/NetflixApp/netflix-reuse-official_ntcjl2_wtrhhh.jpg'/>
      
    </BackgroundContainer>
  )
}

const BackgroundContainer = styled.div`
// we have to write all the css here only
  height: 100vh;
  width: 100vw;
  img{
    height: 100vh;
    width: 100vw;
  }
`

export default BackgroundImage
