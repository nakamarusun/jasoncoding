import React, { useCallback } from "react";
import { useGoogleReCaptcha } from "react-google-recaptcha-v3";

const ContactInfo = () => {
  const { executeRecaptcha } = useGoogleReCaptcha();

  const handleReCaptchaVerify = useCallback(async () => {
    if (!executeRecaptcha) return;

    const token = await executeRecaptcha("getcontact");

    const res = await (
      await fetch("http://localhost:5999/getcontact", {
        method: "POST",
        headers: {
          Accept: "application/json",
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          response: token,
        }),
      })
    ).json();

    console.log(res);
  }, [executeRecaptcha]);

  return <button onClick={handleReCaptchaVerify}>get</button>;
};

export default ContactInfo;
