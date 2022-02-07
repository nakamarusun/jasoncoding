import React from "react";
import PropTypes from "prop-types";
import { techs } from "../../technologies";

// TODO: LazyLoad, animate
const Skills = ({ className }) => {
  return (
    <div className={className}>
      {techs.map(({ img, name }, i) => {
        const InSvg = img;
        return (
          <div className="group" key={i}>
            <InSvg className="h-8 w-8 transition-transform duration-100 hover:scale-125" />
            <div className="bg-white -translate-x-[25%] px-2 py-0.5 rounded-md -translate-y-16 absolute hidden group-hover:block">
              {name}
            </div>
          </div>
        );
      })}
    </div>
  );
};

Skills.propTypes = {
  className: PropTypes.string,
};

export default Skills;
