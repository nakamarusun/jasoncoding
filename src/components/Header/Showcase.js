import React, { useEffect, useRef } from "react";
import Logo from "../../images/logo.inline.svg";
import anime from "animejs";
import PropTypes from "prop-types";

const Showcase = ({ skipAnim, setSkipAnim }) => {
  // We have to use a reference so that it's not lost in the anime context
  const skipAnimRef = useRef(skipAnim);

  // This function triggers the sliding animation for informations
  function slideInfo() {
    anime
      .timeline({})
      .add({
        targets: "#idpanel",
        width: ["0%", "100%"],
        height: ["0%", "100%"],
        easing: "easeInOutQuart",
        duration: 500,
      })
      .add(
        {
          targets: "#skipanimbtn",
          opacity: ["1.0", "0.0"],
          duration: 500,
          complete: () => {
            document.querySelector("#skipanimbtn").style.display = "none";
          },
        },
        0
      );
  }

  useEffect(() => {
    skipAnimRef.current = skipAnim;
    if (skipAnim) {
      slideInfo(true);
    }
  }, [skipAnim]);

  useEffect(() => {
    // Recolor path to white
    const $logoPath = document.querySelectorAll("#mainlogo path");
    $logoPath.forEach((el) => {
      el.style.stroke = "white";
    });

    anime
      // Animation for main title card
      .timeline({
        targets: $logoPath,
      })
      .add({
        strokeDashoffset: [anime.setDashoffset, 0],
        easing: "easeInOutSine",
        delay: anime.stagger(300),
      })
      .add(
        {
          fill: ["rgba(0, 0, 0, 0)", "rgba(255, 255, 255, 1)"],
          easing: "easeOutBounce",
          duration: 500,
          delay: anime.stagger(20, { from: "center" }),
        },
        "+=400"
      )
      .add(
        {
          // Animation for role text
          targets: "#roletext",
          opacity: ["0.0", "1.0"],
          duration: 1,
        },
        "+=800"
      )
      .add(
        {
          targets: ".showcasebutton",
          opacity: ["0.0", "1.0"],
          duration: 1,
          complete: () => {
            // When the animation is completed, check for skip anim.
            if (!skipAnimRef.current) setTimeout(slideInfo, 400);
          },
        },
        "+=800"
      );
  }, []);

  return (
    <div className="relative w-full h-full flex flex-row justify-center bg-black">
      <h5
        id="skipanimbtn"
        onClick={setSkipAnim}
        className="absolute text-white font-bold bottom-4 left-4 text-center underline cursor-pointer select-none"
      >
        Skip animation.
      </h5>
      <div className="py-4 max-w-3xl w-11/12 flex flex-col justify-center items-center">
        <Logo
          // // className="max-h-[50%] h-full" // Forgive me, for i have to commit some evilness.
          id="mainlogo"
          className="max-h-full"
        />
        <h2
          id="roletext"
          className="leading-none md:leading-none opacity-0 my-3 font-serif text-white w-full text-right text-md md:text-lg space origin-right scale-y-[175%]"
        >
          Computer Science Student, Full-Stack Developer, Designer
        </h2>
        <div className="opacity-0 showcasebutton material-icons text-white animate-bounce mt-[1rem] mb-[-.25rem]">
          adjust
        </div>
        <button
          style={{
            transition: "border-radius 0.5s cubic-bezier(0.22, 1, 0.36, 1)",
          }}
          className="opacity-0 showcasebutton font-serif border-2 border-gray-50 px-6 pt-1 pb-2 text-md md:text-xl font-bold text-white rounded-3xl hover:rounded-lg scale-x-[80%] bg-gradient-to-b from-black to-gray-900"
        >
          PROJECT SHOWCASE
        </button>
      </div>
    </div>
  );
};

Showcase.propTypes = {
  skipAnim: PropTypes.bool.isRequired,
  setSkipAnim: PropTypes.func.isRequired,
};

export default Showcase;
