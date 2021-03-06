#ifndef OSS_CLASSIC_UI_TOOLBOX_WINDOW_H
#define OSS_CLASSIC_UI_TOOLBOX_WINDOW_H

#include "../game.h"


namespace OSS {
	namespace Classic {	
		class ToolsGroupButton;
		class Button;
		
		class ToolsWindow : public Window, public Responder {
			
			/**
			 * Construction
			 */
		public:
			const Pointer<GameUI> ui;
			
			ToolsWindow(GameUI * ui);
			
			//TODO: move this somewhere else
			Tower * getTower();
			
			
			/**
			 * Subviews
			 */
		private:
			typedef map< ItemGroup, Pointer<ToolsGroupButton> > GroupButtonMap;
			GroupButtonMap groupButtons;
			
			Pointer<Button> bulldozerButton;
			Pointer<Button> fingerButton;
			Pointer<Button> inspectorButton;
			
			
			/**
			 * State
			 */
		public:
			virtual void update();
			virtual void updateButtons();
			virtual void layout();
			
			Updatable::Conditional<ToolsWindow> updateButtonsIfNeeded;
			Updatable::Conditional<ToolsWindow> layoutIfNeeded;
			
			
			/**
			 * Drawing
			 */			
		public:
			virtual void draw(rectd dirtyRect);
			
			
			/**
			 * Events
			 */
		public:
			virtual void eventToolChanged(Event * event);
		};
	}
}


#include "groupbutton.h"
#include "../button.h"


#endif
