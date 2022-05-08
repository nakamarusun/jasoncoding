import React, { useCallback, useState } from "react";
import { useGoogleReCaptcha } from "react-google-recaptcha-v3";
import PropTypes from "prop-types";
import { API_ROOT } from "../../constants";

const ContactInfo = ({ className }) => {
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
    try {
      const res = await (
        await fetch(API_ROOT + "/getcontact", {
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

      setContact({
        status: "success",
        data: res,
      });
    } catch (e) {
      console.error(e);
      setContact({
        ...contact,
        status: "error",
      });
    }
  }, [executeRecaptcha, showTos]);

  const contactStatus = contact.status;

  return (
    <div
      className={`w-full mb-4 flex flex-col items-center ${
        className ? className : ""
      }`}
    >
      {contactStatus === "idle" ? (
        <>
          <button
            className="transition-colors duration-75 text-center px-3 py-1 rounded-md text-purple-900 shadow-nm-sm bg-gray-200 hover:bg-nm-up active:bg-nm-down hover:text-fuchsia-700"
            onClick={handleReCaptchaVerify}
          >
            <span className="material-icons align-bottom mr-2">
              contact_phone
            </span>
            <span className="font-bold ">Get my contact</span>
          </button>
          {showTos && (
            <p className="text-xs w-72 text-center mt-1">
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
        <div className="shadow-nm-sm-inset bg-gray-200 rounded-lg px-4 py-2 flex flex-row items-center">
          {(() => {
            switch (contactStatus) {
              case "getting":
                return (
                  <>
                    <div className="material-icons mr-3 text-purple-900 animate-spin">
                      autorenew
                    </div>
                    <h3>
                      <span className="font-medium text-gray-700">
                        Getting Contact
                      </span>
                    </h3>
                  </>
                );
              case "success":
                return (
                  <>
                    <div className="material-icons mr-3 text-purple-900">
                      contact_phone
                    </div>
                    <div>
                      {Object.entries(contact.data).map(([key, val], i) => (
                        <h3 key={i}>
                          <span className="font-medium text-gray-700">
                            {key.charAt(0).toUpperCase() + key.slice(1)}:&nbsp;
                          </span>
                          <span className="select-all">{val}</span>
                        </h3>
                      ))}
                    </div>
                  </>
                );
              case "error":
                return (
                  <>
                    <div className="material-icons mr-2 text-purple-900">
                      error
                    </div>
                    <h3>
                      <span className="font-medium text-red-600">
                        Captcha Failed
                      </span>
                    </h3>
                  </>
                );
            }
          })()}
        </div>
      )}
    </div>
  );
};

ContactInfo.propTypes = {
  className: PropTypes.string,
};

export default ContactInfo;
