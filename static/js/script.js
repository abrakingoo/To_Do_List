document.addEventListener('DOMContentLoaded', function() {
  const triggers = document.querySelectorAll('.trigger');
  const hoverElements = document.querySelectorAll('.hover-element');
  
  // Attach events to each trigger
  triggers.forEach((trigger, index) => {
    const hoverElement = hoverElements[index];

    // Show hover element on trigger hover
    trigger.addEventListener('mouseover', function(event) {
      hoverElement.style.display = 'block';
      updateHoverPosition(event, hoverElement);
    });

    // Hide hover element when cursor leaves trigger
    trigger.addEventListener('mouseleave', function() {
      hoverElement.style.display = 'none';
    });

    // Update hover element position on mousemove
    trigger.addEventListener('mousemove', function(event) {
      updateHoverPosition(event, hoverElement);
    });
  });

  function updateHoverPosition(event, hoverElement) {
    const xOffset = 10; // Adjust as needed to offset from cursor
    const yOffset = 10; // Adjust as needed to offset from cursor

    hoverElement.style.left = (event.pageX + xOffset) + 'px';
    hoverElement.style.top = (event.pageY + yOffset) + 'px';
  }
});

