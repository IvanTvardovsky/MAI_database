import React from "react";

class Image extends React.Component {
   render() {
      return (
         <img className="rounded-image" src={this.props.image} alt=''/>
      )
   }
}

export default Image