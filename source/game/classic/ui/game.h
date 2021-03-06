#ifndef OSS_CLASSIC_UI_GAME_H
#define OSS_CLASSIC_UI_GAME_H

#include "../game.h"

#include "../../ui/view.h"
#include "../../ui/window.h"


namespace OSS {
	namespace Classic {
		class ToolsUI;
		
		class ControlWindow;
		class ToolsWindow;
		
		class GameUI : public GameObject {
			
			/**
			 * Construction
			 */
		public:
			const Pointer<GameScene> scene;
			
			GameUI(GameScene * scene);
			
			//TODO: move this somewhere else
			Tower * getTower();
			
			
			/**
			 * Subsystems
			 */
		public:
			Pointer<View> rootView;
			Pointer<ToolsUI> tools;
			
			
			/**
			 * Windows
			 */
		private:
			Pointer<ControlWindow> controlWindow;
			//Pointer<MapWindow> mapWindow;
			Pointer<ToolsWindow> toolsWindow;
			
			
			/**
			 * Simulation
			 */
		public:
			virtual void advance(double dt);
			
			
			/**
			 * State
			 */
		public:
			virtual void update();
			virtual void updateRootViewFrame();
			
			Updatable::Conditional<GameUI> updateRootViewFrameIfNeeded;
			
			
			/**
			 * Drawing
			 */
		public:
			virtual void draw(rectd dirtyRect);
			
			
			/**
			 * Events
			 */
		public:
			virtual bool sendEventToNextResponders(OSS::Event * event);
			
			virtual void eventVideoChanged(VideoEvent * event);
		};
	}
}


#include "control/window.h"
#include "toolbox/window.h"

#include "tools/tools.h"


#endif
