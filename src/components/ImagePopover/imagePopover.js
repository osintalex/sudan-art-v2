import {
  Center,
  Image,
  Modal,
  ModalBody,
  ModalCloseButton,
  ModalContent,
  ModalHeader,
  ModalOverlay,
  Tag,
  TagLabel,
  Text,
  Link,
} from "@chakra-ui/react";
import PropTypes from "prop-types";
import React from "react";
import {
  FacebookIcon,
  FacebookShareButton,
  TwitterIcon,
  TwitterShareButton,
  WhatsappIcon,
  WhatsappShareButton,
} from "react-share";
import { makeRepeated } from "../../utils/utils";
import MultiLingualContent from "../MultingualContent/multilingualContent.js";

/**
 * Popover for users to click on images once the search results have returned - this way they can see artist
 * and upload date
 * @param {props} props contains the image details. In addition,  isOpen, onClose, onOpen
 * are all passed down from the Chakra Modal component see - search.js
 * @return {component} a modal popover component
 */
function ImagePopover(props) {
  const tagGradients = makeRepeated(
    [
      "linear(to-r, orange.400, yellow.400)",
      "linear(to-r, teal.400, blue.400)",
      "linear(to-r, pink.400, red.400)",
      "linear(to-r, green.400, teal.400)",
    ],
    5
  );
  const { imageDescription, imageArtist, imageDate, sourceURL } =
    props.popoverImageDetails;

  return (
    <>
      <Modal isOpen={props.isOpen} onClose={props.onClose}>
        <ModalOverlay />
        <ModalContent color="gray.50">
          <ModalHeader className="search-modal-header">
            {imageArtist}
          </ModalHeader>
          <ModalCloseButton />
          <ModalBody className="search-modal-body">
            <Image
              className="search-popover-image"
              alt={imageDescription}
              src={props.popoverImageDetails.imageSrc}
              onClick={props.onOpen}
            />
            {imageDescription.split(",").map((tag, index) => (
              <Tag
                size={"lg"}
                key={`tag ${index}`}
                className="image-tag"
                bgGradient={tagGradients[index]}
              >
                <TagLabel key={`tag label ${index}`} color="gray.50">
                  {tag.toLowerCase()}
                </TagLabel>
              </Tag>
            ))}
            <br />
            <Text fontSize="sm" color="gray.50" as="abbr" align="center">
              {imageDate}
            </Text>
            <br />
            <Link
              fontSize="sm"
              color="gray.50"
              as="abbr"
              align="center"
              href={sourceURL}
              isExternal
            >
              {sourceURL}
            </Link>
            <br />
            <Text fontSize="sm" color="gray.50" as="i" align="center">
              <MultiLingualContent contentID="image_popover" />
            </Text>
            <Center style={{ transform: "scale(0.6)" }}>
              <FacebookShareButton
                url={props.imageSrc}
                quote={
                  "Look at this amazing piece of Sudanese revolutionary art! Burhan fi kobr!"
                }
                hashtag={"#sudancoup"}
              >
                <FacebookIcon />
              </FacebookShareButton>
              <TwitterShareButton
                url={props.imageSrc}
                title={"Sudanese Revolutionary Art"}
                via={"https://sudanart.com"}
                hashtags={["#sudancoup", "#sudanart"]}
                related={["bsonblast"]}
              >
                <TwitterIcon />
              </TwitterShareButton>
              <WhatsappShareButton
                url={props.imageSrc}
                title={"Check out this amazing Sudanese art!"}
              >
                <WhatsappIcon />
              </WhatsappShareButton>
            </Center>
          </ModalBody>
        </ModalContent>
      </Modal>
    </>
  );
}
ImagePopover.propTypes = {
  popoverImageDetails: PropTypes.object,
  isOpen: PropTypes.bool,
  onOpen: PropTypes.bool,
  onClose: PropTypes.func,
};
export default ImagePopover;
