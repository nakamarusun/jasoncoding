import React, { useCallback, useState } from "react";
import { useGoogleReCaptcha } from "react-google-recaptcha-v3";

const ContactInfo = () => {
  // Google recaptcha hook
  const { executeRecaptcha } = useGoogleReCaptcha();

  // Whether the contact has been accepted. 0 Is before pressing. 1 is
  // confirming the TOS. 2 is get data.
  const [showTos, setshowTos] = useState(false);
  const [contact, setContact] = useState({
    status: "idle",
    data: {},
  });

  // Sends the captcha data and gets the contact data.
  const handleReCaptchaVerify = useCallback(async () => {
    if (!showTos) {
      setshowTos(true);
      return;
    }

    if (!executeRecaptcha) return;

    // Set status to loading
    setContact({
      ...contact,
      status: "getting",
    });

    // Execute the captcha with the getcontact action
    const token = await executeRecaptcha("getcontact");

    // Asks the backend for our contact
    const res = await (
      await fetch("https://jasoncoding.com/api/getcontact", {
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

    setContact({
      status: "success",
      data: res,
    });
  }, [executeRecaptcha, showTos]);

  const contactStatus = contact.status;

  return (
    <div className="w-full mb-4 flex flex-col items-center">
      {contactStatus === "idle" ? (
        <>
          <button
            className="text-center border border-gray-600 px-3 py-1 rounded-md text-purple-900 hover:bg-gray-300 hover:text-fuchsia-700"
            onClick={handleReCaptchaVerify}
          >
            <span className="material-icons align-bottom mr-2">
              contact_phone
            </span>
            <span className="font-bold ">Get my contact</span>
          </button>
          {showTos && (
            <p className="text-xs w-72 text-center">
              This site is protected by reCAPTCHA and the Google&nbsp;
              <a
                className="underline text-blue-600"
                href="https://policies.google.com/privacy"
              >
                Privacy Policy&nbsp;
              </a>
              and&nbsp;
              <a
                className="underline text-blue-600"
                href="https://policies.google.com/terms"
              >
                Terms of Service&nbsp;
              </a>
              apply. To get my contact, click the button again.
            </p>
          )}
        </>
      ) : (
        <div className="bg-white rounded-lg px-3 py-0.5 flex flex-row items-center">
          <div
            className={`material-icons mr-2 text-purple-900 ${
              contactStatus === "getting" ? "animate-spin" : ""
            }`}
          >
            {contactStatus === "getting" ? "autorenew" : "lock_open"}
          </div>
          <div>
            <h3>
              <span className="font-medium text-gray-700">Email:&nbsp;</span>
              {contact.data.email}
            </h3>
            <h3>
              <span className="font-medium text-gray-700">Phone:&nbsp;</span>
              {contact.data.phone}
            </h3>
          </div>
        </div>
      )}
    </div>
  );
};

export default ContactInfo;
