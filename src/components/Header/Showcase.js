import React, { useEffect } from "react";
import Logo from "../../images/logo.inline.svg";
import anime from "animejs";

const Showcase = () => {
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
        },
        "+=800"
      )
      .add(
        {
          targets: "#idpanel",
          width: ["0%", "100%"],
          easing: "easeInOutQuart",
          duration: 500,
        },
        "+=800"
      );
  }, []);

  return (
    <div className="w-full flex flex-row justify-center bg-black">
      <div className="max-w-3xl w-11/12 h-full flex flex-col justify-center items-center">
        <Logo
          // // className="max-h-[50%] h-full" // Forgive me, for i have to commit some evilness.
          id="mainlogo"
          className="max-h-full h-auto"
        />
        <h2
          id="roletext"
          className="opacity-0 mt-1 font-serif text-white text-right w-full text-4xl origin-right scale-x-[60%] "
        >
          {/* COMPUTER SCIENCE STUDENT, FULL-STACK DEVELOPER, DESIGNER */}
          Computer Science Student, Full-Stack Developer, Designer
        </h2>
        <div className="opacity-0 showcasebutton material-icons text-white animate-bounce mt-[1rem] mb-[-.25rem]">
          adjust
        </div>
        <button
          style={{
            transition: "border-radius 0.5s cubic-bezier(0.22, 1, 0.36, 1)",
          }}
          className="opacity-0 showcasebutton font-serif border-2 border-gray-50 px-6 pt-1 pb-2 text-xl font-bold text-white rounded-3xl hover:rounded-lg scale-x-[80%] bg-gradient-to-b from-black to-gray-900"
        >
          PROJECT SHOWCASE
        </button>
      </div>
    </div>
  );
};

export default Showcase;
