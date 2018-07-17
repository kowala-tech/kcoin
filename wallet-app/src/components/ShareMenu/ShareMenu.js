import React from "react";
import PropTypes from "prop-types";
import { compose } from "recompose";
// Material UI Imports
import { withStyles } from "@material-ui/core/styles";
// Component Related Imports
import styles from "./styles";
import {
	FacebookShareButton,
	GooglePlusShareButton,
	LinkedinShareButton,
	TwitterShareButton,
	TelegramShareButton,
	WhatsappShareButton,
	RedditShareButton,
	EmailShareButton,

	FacebookIcon,
	TwitterIcon,
	GooglePlusIcon,
	LinkedinIcon,
	TelegramIcon,
	WhatsappIcon,
	RedditIcon,
	EmailIcon,
} from "react-share";

class ShareMenu extends React.Component {
	render() {
		const { classes, text } = this.props;
		const shareUrl = "https://www.kowala.tech";
		const title = `My kUSD wallet address is ${text}`;

		return (
			<div className={classes.container}>
				<div className={classes.buttonContainer}>
					<FacebookShareButton
						url={shareUrl}
						quote={title}
						className={classes.button}>
						<FacebookIcon
							size={32}
							round />
					</FacebookShareButton>
				</div>

				<div className={classes.buttonContainer}>
					<TwitterShareButton
						url={shareUrl}
						title={title}
						className={classes.button}>
						<TwitterIcon
							size={32}
							round />
					</TwitterShareButton>
				</div>

				<div className={classes.buttonContainer}>
					<TelegramShareButton
						url={shareUrl}
						title={title}
						className={classes.button}>
						<TelegramIcon size={32}
							round />
					</TelegramShareButton>
				</div>

				<div className={classes.buttonContainer}>
					<WhatsappShareButton
						url={shareUrl}
						title={title}
						separator=":: "
						className={classes.button}>
						<WhatsappIcon size={32}
							round />
					</WhatsappShareButton>
				</div>

				<div className={classes.buttonContainer}>
					<GooglePlusShareButton
						url={shareUrl}
						className={classes.button}>
						<GooglePlusIcon
							size={32}
							round />
					</GooglePlusShareButton>
				</div>

				<div className={classes.buttonContainer}>
					<LinkedinShareButton
						url={shareUrl}
						title={title}
						windowWidth={750}
						windowHeight={600}
						className={classes.button}>
						<LinkedinIcon
							size={32}
							round />
					</LinkedinShareButton>
				</div>

				<div className={classes.buttonContainer}>
					<RedditShareButton
						url={shareUrl}
						title={title}
						windowWidth={660}
						windowHeight={460}
						className={classes.button}>
						<RedditIcon
							size={32}
							round />
					</RedditShareButton>
				</div>

				<div className={classes.buttonContainer}>
					<EmailShareButton
						url={shareUrl}
						subject={title}
						body="body"
						className={classes.button}>
						<EmailIcon
							size={32}
							round />
					</EmailShareButton>
				</div>
			</div>
		);
	}
}

ShareMenu.propTypes = {
	classes: PropTypes.object.isRequired,
	text: PropTypes.string.isRequired
};

export default compose(
	withStyles(styles)
)(ShareMenu);
