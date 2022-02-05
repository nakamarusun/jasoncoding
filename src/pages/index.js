import * as React from "react";
import Header from "../components/Header/Header";
import { Helmet } from "react-helmet";
import { GoogleReCaptchaProvider } from "react-google-recaptcha-v3";
import Sidebar from "../components/Sidebar";

const IndexPage = () => {
  return (
    <GoogleReCaptchaProvider
      reCaptchaKey="6LdGuFkeAAAAAFJRKFLPVB5Cd51jY0R1GKpsCZnL"
      scriptProps={{ defer: true, async: true }}
    >
      <Helmet titleTemplate="Jasoncoding - %s" defaultTitle="Jasoncoding">
        <html lang="en" amp />
        <meta charSet="utf-8" />
        <link rel="canonical" href="https://jasoncoding.com" />
        <meta
          name="description"
          content="Jasoncoding. Jason Christian Hailianto programming and development portfolio website"
        />
      </Helmet>
      <Sidebar />
      <Header />
    </GoogleReCaptchaProvider>
  );
};

export default IndexPage;
